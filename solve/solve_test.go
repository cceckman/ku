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

	caseTwo = `Case 2
1324
2413
4231
3142`

	caseThree = `Case 3
870000000300D040
05E0003AF9100600
G00058700E0090BC
00400E6D0BAC0703
E0080010000374C0
900006FC00DE0310
B0A300D008010060
6001E0400509B00D
0000F0000090502A
A1006050DF7G0000
00GB040820000D07
09071302608AGFE4
70D09G0500E43802
0030A00000000GF0
109000E400000070
06803000A7000000
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

	if solved, issues := verify.IsSolved(p); !solved {
		t.Errorf("Puzzle %s isn't solved: \n%s\n", p.Name, strings.Join(issues, "\n\t"))
	}

	out := bytes.NewBuffer(make([]byte,0))
	p.WriteTo(out)
	t.Logf("Solution:\n")
	t.Logf(out.String())
}

func TestSolveTwo(t *testing.T) {
	r := strings.NewReader(caseTwo)

	p, err := puzzle.NewPuzzle(2, r)
	if err != nil {
		t.Fatalf("Error instantiating test case: %v", err)
	}
	if err := Solve(p); err != nil {
		t.Fatalf("Error from solver: %v", err)
	}

	if solved, issues := verify.IsSolved(p); !solved {
		t.Errorf("Puzzle %s isn't solved: \n%s\n", p.Name, strings.Join(issues, "\n\t"))
	}

	out := bytes.NewBuffer(make([]byte,0))
	p.WriteTo(out)
	t.Logf("Solution:\n")
	t.Logf(out.String())
}

func TestSolveLarge(t *testing.T) {
	r := strings.NewReader(caseThree)

	p, err := puzzle.NewPuzzle(4, r)
	if err != nil {
		t.Fatalf("Error instantiating test case: %v", err)
	}
	if err := Solve(p); err != nil {
		t.Fatalf("Error from solver: %v", err)
	}

	if solved, issues := verify.IsSolved(p); !solved {
		t.Errorf("Puzzle %s isn't solved: \n%s\n", p.Name, strings.Join(issues, "\n\t"))
	}

	out := bytes.NewBuffer(make([]byte,0))
	p.WriteTo(out)
	t.Logf("Solution:\n")
	t.Logf(out.String())
}
