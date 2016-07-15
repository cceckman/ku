package solve

import (
	"fmt"
	"github.com/cceckman/ku/puzzle"
)

func Solve(p *puzzle.Puzzle) error {
	// Start super simple (serial).
	// Continue as long as we had some progress; at most 
	for i, changed := 0, true; changed; i, changed = i+1, false {
		fmt.Printf("Starting iteration %d\n", i)
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
					fmt.Printf("Only one option for %s\n", p.CellInfo(i))
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
						fmt.Printf("Only one place for %s\n", p.CellInfo(list[0]))
						changed = true
					} else if len(list) == 1 && p.Value[list[0]] != v {
						return fmt.Errorf("inconsistent result: cell %s has value %d from solver, but already has value %d",
							p.CellInfo(list[0]), v, p.Value[list[0]])
					}
				}
			}
		}
	}
	return nil
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
	allValues = append(allValues, p.Values(p.Box(p.BoxOf(idx)))...)

	// At the beginning, we said that we'd exit early if this cell was already filled.
	for _, v := range allValues {
		if v != 0 {
			delete(possibilities, v)
		}
	}
	return possibilities
}


