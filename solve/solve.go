package solve

import (
	"github.com/cceckman/ku/puzzle"
)

func Solve(p *puzzle.Puzzle) *puzzle.Puzzle {
	// Start super simple (serial).
	// Continue as long as we had some progress; at most 
	for changed := true; changed; {
		// Step 1: Iterate over cells. Anything that can only be one thing, must be that thing.
		for i, v := range p.Value {
			if v != 0 {
				continue // Already solved for.
			}
			possible := possible(p, i)
			if len(possible) == 1 {
				for k, _ := range possible {
					p.Value[i] = k
					changed = true
				}
			}
		}

		// Step 2: Iterate over zones. Any value that only can be one cell, must be.
		for i := 0; i < p.Size * p.Size; i++ {
			groups := [][]int{p.Row(i), p.Col(i), p.Box(i)}
			for _, group := range groups {
				for v, list := range valuePossibilities(p, group) {
					if len(list) == 1 && p.Value[list[0]] == 0 {
						// Found only one cell that could solve this.
						p.Value[list[0]] = v
						changed = true
					} else if len(list) == 1 && p.Value[list[0]] != v {
						panic("Inconsistent result!")
					}
				}
			}
		}
	}
	return p
}

// Solve for, within the given cells, what cell each value could be in.
// Result includes existing values.
func valuePossibilities(p *puzzle.Puzzle, idx []int) map[uint64][]int {
	possibilities := make(map[uint64][]int)
	for _, i := range idx {
		cellPossible := possible(p, i)
		for k := range cellPossible {
			possibilities[k] = append(possibilities[k], i)
		}
	}
	return possibilities
}

// Solve for what this cell could be, based on its row, column, and box.
func possible(p *puzzle.Puzzle, idx int) map[uint64]bool {
	possibilities := make(map[uint64]bool)
	if p.Value[idx] != 0 {
		possibilities[p.Value[idx]] = true
		return possibilities
	}

	for i := 1; i < p.Size * p.Size + 1; i++ {
		possibilities[uint64(i)] = true
	}
	allValues := make([]uint64, 0, p.Size * p.Size * 3)
	allValues = append(allValues, p.Values(p.Row(p.RowOf(idx)))...)
	allValues = append(allValues, p.Values(p.Col(p.ColOf(idx)))...)
	allValues = append(allValues, p.Values(p.Box(p.RowOf(idx)))...)

	for _, v := range allValues {
		if v != 0 {
			delete(possibilities, v)
		}
	}
	return possibilities
}


