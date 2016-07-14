package sudoku

import (
	"strings"
	"testing"
)

const (
	firstCase = `Case 1
000000001
603091005
079040080
050074000
000002006
000030000
504060090
006008004
300000700`
	secondCase = `Case 4
198734265
564192378
273865914
315427689
849356721
627918543
736541892
452689137
981273456`
	firstName = "Case 1"
	secondName = "Case 4"
	// TODO add >3x3 case
)

func TestSinglePuzzle(t *testing.T) {
	firstReader := strings.NewReader(firstCase)
	p, err := NewPuzzle(3, firstReader)
	if err != nil {
		t.Errorf("got error when loading first puzzle: %v", err)
	}
	if p.name != firstName {
		t.Errorf("puzzle name doesn't match. got: %v expected: %v", p.name, firstName)
	}

	// Spot-check RowOf, ColOf, BoxOf; test case is index, row, column, box.
	for n, test := range [][]int{
		[]int{0, 0, 0, 0},
		[]int{1, 0, 1, 0},
		[]int{40, 4, 4, 4},
		[]int{35, 3, 8, 5},
	} {
		if p.RowOf(test[0]) != test[1] {
			t.Errorf("RowOf test %d failed: got: %v expected: %v", n, p.RowOf(test[0]), test[1])
		}
		if p.ColOf(test[0]) != test[2] {
			t.Errorf("ColOf test %d failed: got: %v expected: %v", n, p.ColOf(test[0]), test[2])
		}
		if p.BoxOf(test[0]) != test[3] {
			t.Errorf("BoxOf test %d failed: got: %v expected: %v", n, p.RowOf(test[0]), test[3])
		}
	}
}
