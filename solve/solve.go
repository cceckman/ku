package solve

import (
	"fmt"
	"github.com/cceckman/ku/puzzle"
	"github.com/cceckman/ku/verify"
)

func Solve(p *puzzle.Puzzle) error {
	// Start super simple (serial).
	// Continue as long as we had some progress.
	var iteration int
	for iteration, changed := 0, true; changed; iteration, changed = iteration+1, false {
		fmt.Printf("Starting iteration %d\n", iteration) // DEBUG
		// Step 1: Iterate over cells. Anything that can only be one thing, must be that thing.
		for i, v := range p.Value {
			if v != 0 {
				continue // Already solved for.
			}
			possible := possible(p, i)
			if len(possible) == 0 {
				return fmt.Errorf("No logical possibilities for %s, stopping progress", p.CellInfo(i))
			}
			if len(possible) == 1 {
				for k, _ := range possible {
					p.Value[i] = k
					changed = true
					fmt.Printf("Only one option for %s\n", p.CellInfo(i)) // DEBUG
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
						fmt.Printf("Only one place for %s\n", p.CellInfo(list[0])) // DEBUG
						changed = true
					} else if len(list) == 1 && p.Value[list[0]] != v {
						return fmt.Errorf("inconsistent result: cell %s has value %d from solver, but already has value %d",
							p.CellInfo(list[0]), v, p.Value[list[0]])
					}
				}
			}
		}
	} // end "without guessing" stanza.

	if solved, _ := verify.IsSolved(p); solved {
		return nil // Solved! Move along.
	}
	// Not yet solved. Do a DFS across all remaining possibilities.
	fmt.Printf("Haven't solved after %d iterations; starting to guess.\n", iteration) // DEBUG
	for i, v := range p.Value {
		if v != 0 {
			// Cell already filled deterministically.
			continue
		}
		options := possible(p, i)
		for option, _ := range options {
			// Guess option for i in a new puzzle.
			guessed := p.Copy()
			guessed.Value[i] = option
			fmt.Printf("Guessing: %s\n", guessed.CellInfo(i))
			err := Solve(guessed)
			if err == nil {
				// A solution was found along that branch.
				fmt.Printf("Solution found by guessing: %s\n", guessed.CellInfo(i))
				// Copy result back up
				p.Value = guessed.Value
				return nil
			} //else: branch was discarded, move on to the next one.
		}
	}
	return fmt.Errorf("No branch found a solution.")
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


