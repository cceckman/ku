package verify

import (
	"fmt"

	"github.com/cceckman/ku/puzzle"
)

// Checks if a Puzzle is solved and is a valid solution.
// Returns a possibly-empty list of reasons why it isn't a valid solution.
// This is not designed to be efficient.
func IsSolved(p *puzzle.Puzzle) (bool, []string) {
	sizeSq := p.Size * p.Size

	var issues []string
	for n := 0; n < sizeSq; n++ {
		// Validate row, column, and box n.
		for _, issue := range complete(p, p.Row(n)) {
			issues = append(issues, fmt.Sprintf("Row %d invalid: %s", n, issue))
		}
		for _, issue := range complete(p, p.Col(n)) {
			issues = append(issues, fmt.Sprintf("Col %d invalid: %s", n, issue))
		}
		for _, issue := range complete(p, p.Box(n)) {
			issues = append(issues, fmt.Sprintf("Box %d invalid: %s", n, issue))
		}
	}

	return (len(issues) == 0), issues
}

// complete checks if the indices in the given puzzle are a complete set,
// e.g. have exactly the necessary values.
func complete(p *puzzle.Puzzle, idx []int) []string {
	// Maps int to where it's found.
	mask := make([]int, p.Size*p.Size+1) // Indexed range one to maximum.
	for i := range mask {
		mask[i] = -1 // 0 is a valid cell index! So, preset to something that isn't a valid index.
		// Panic! At The Index, if it's actually used as an index. Which of course it shouldn't be,
		// but didd the first time I ran this, because I hadn't updated the "is this index already seen" line.
	}

	var issues []string

	for _, i := range idx {
		v := p.Value[i]
		if v == 0 {
			issues = append(issues,
				fmt.Sprintf("puzzle not solved: index %d isn't set (%s)", i, p.CellInfo(i)))
		}
		if mask[v] != -1 {
			issues = append(issues,
				fmt.Sprintf("value %d found at two indices: (%s) (%s)", v, p.CellInfo(i), p.CellInfo(mask[v])))
		}
		// Add to mask, to indicate it's present.
		mask[v] = i
	}

	// The "no duplicates, no zeros" should check this, but let's make sure...
	for v := 1; v < len(mask); v++ {
		if mask[v] == -1 {
			issues = append(issues, fmt.Sprintf("value %d not found", v))
		}
	}
	return issues
}
