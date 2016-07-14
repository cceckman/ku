package verify

import (
	"github.com/cceckman/ku/puzzle"
)

// Checks if a Puzzle is solved and is a valid solution.
// Returns (true, "") if it is, (false, why not) if it isn't.
func IsSolved(p *puzzle.Puzzle) (bool, string) {
	sizeSq := p.Size * p.Size

	// A mask of "is this value found."
	// TODO: Measure whether this is more or less efficient than a map
	_ = make([]bool, sizeSq)

	return false, ""
}
