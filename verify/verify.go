package verify

import (
	"fmt"

	"github.com/cceckman/ku/puzzle"
)

// Checks if a Puzzle is solved and is a valid solution.
// Returns (true, "") if it is, (false, why not) if it isn't.
// This is not designed to be efficient.
func IsSolved(p *puzzle.Puzzle) (bool, []string) {
	sizeSq := p.Size * p.Size

	var issues []string
	for n := 0; n < sizeSq; n++ {
		// Validate row, column, and box n.
		for _, issue := range complete(p, p.Row(n)) {
			issues = append(issues, fmt.Sprintf("Row %d invalid: %s", issue))
		}
		for _, issue := range complete(p, p.Col(n)) {
			issues = append(issues, fmt.Sprintf("Col %d invalid: %s", issue))
		}
		for _, issue := range complete(p, p.Box(n)) {
			issues = append(issues, fmt.Sprintf("Box %d invalid: %s", issue))
		}
	}

	return (len(issues) == 0), issues
}

// complete checks if the indices in the given puzzle are a complete set,
// e.g. have exactly the necessary values.
func complete(p *puzzle.Puzzle, idx []int) []string {
	// Maps int to where it's found.
	mask := make([]int, p.Size*p.Size+1) // +1 because 0 should not be found.

	var issues []string

	for _, i := range idx {
		v := p.Value[i]
		if v == 0 {
			issues = append(issues,
				fmt.Sprintf("puzzle not solved: index %d isn't set (%s)", i, p.CellInfo(i)))
		}
		if mask[v] != 0 {
			issues = append(issues,
				fmt.Sprintf("value %d found at two indices: (%s) (%s)", v, p.CellInfo(i), p.CellInfo(mask[v])))
		}
	}

	// The "no duplicates, no zeros" should check this, but let's make sure...
	for v := 1; v < len(mask); v++ {
		if mask[v] == 0 {
			issues = append(issues, fmt.Sprintf("value %d not found", v))
		}
	}
	return issues
}
