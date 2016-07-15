package solve

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cceckman/ku/puzzle"
	"github.com/cceckman/ku/verify"
)

const (
	caseOne = `Case 1
007030100
340000026
809006703
003708090
000000000
050904600
702300405
530000018
006040300
`
)

func TestSolve(t *testing.T) {
	r := strings.NewReader(caseOne)

	p, err := puzzle.NewPuzzle(3, r)
	if err != nil {
		t.Fatalf("Error instantiating test case: %v", err)
	}
	if err := Solve(p); err != nil {
		t.Fatalf("Error from solver: %v", err)
	}

	if solved, issues := verify.IsSolved(p); solved {
		t.Errorf("Puzzle %s isn't solved: \n%s\n", p.Name, strings.Join(issues, "\n\t"))
	}

	out := bytes.NewBuffer(make([]byte,0))
	p.Print(out)
	t.Errorf("Solution:\n")
	t.Errorf(out.String())
}
