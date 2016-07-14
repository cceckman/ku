package puzzle

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

const (
	tradPrefix = "3 2\n"
	firstCase  = `Case 1
000000001
603091005
079040080
050074000
000002006
000030000
504060090
006008004
300000700
`
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
	firstName  = "Case 1"
	secondName = "Case 4"
	// TODO add >3x3 case
)

func TestCollection(t *testing.T) {
	// TODO test invalid collections, e.g. 2 of the same name
	prefixReader := strings.NewReader(tradPrefix)
	firstReader := strings.NewReader(firstCase)
	secondReader := strings.NewReader(secondCase)
	catReader := io.MultiReader(prefixReader, firstReader, secondReader)

	expectedOutput := new(bytes.Buffer)
	r := io.TeeReader(catReader, expectedOutput)

	collection, err := NewCollection(r)
	if err != nil {
		t.Fatalf("couldn't create Collection: %v", err)
	}

	// Test "print"; should match the input read.
	output := new(bytes.Buffer)
	collection.Print(output)

	// NB: Print always terminates with a newline, but it doesn't care whether there's a trailing newline.
	expectedOutput.WriteString("\n")
	if output.String() != expectedOutput.String() {
		t.Errorf("Print failed:\ngot:\n%v\nexpected:\n%v\n---\n", output, expectedOutput.String())
	}
}

func TestSinglePuzzle(t *testing.T) {
	firstReader := strings.NewReader(firstCase)
	p, err := NewPuzzle(3, firstReader)
	if err != nil {
		t.Errorf("got error when loading first puzzle: %v", err)
	}
	if p.Name != firstName {
		t.Errorf("puzzle name doesn't match. got: %v expected: %v", p.Name, firstName)
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

	// Test "print"; should match input.
	output := bytes.NewBuffer(make([]byte, 0, len(firstCase)))
	p.Print(output)

	// NB: Print always terminates with a newline, but it doesn't care whether there's a trailing newline.
	if output.String() != firstCase {
		t.Errorf("Print failed:\ngot:\n%v\nexpected:\n%v\n---\n", output, firstCase)
	}
}

// TestTwoPuzzles confirms that reading one puzzle after another, from the same reader, works correctly.
// This also gets tested as part of collection_test, but test it here explicitly.
func TestTwoPuzzles(t *testing.T) {
	firstReader := strings.NewReader(firstCase)
	secondReader := strings.NewReader(secondCase)
	r := io.MultiReader(firstReader, secondReader)

	firstPuzzle, err := NewPuzzle(3, r)
	if err != nil {
		t.Errorf("got error when loading first puzzle: %v", err)
	}
	secondPuzzle, err := NewPuzzle(3, r)
	if err != nil {
		t.Errorf("got error when loading second puzzle: %v", err)
	}
	if firstPuzzle.Name != firstName {
		t.Errorf("puzzle name doesn't match. got: %v expected: %v", firstPuzzle.Name, firstName)
	}
	if secondPuzzle.Name != secondName {
		t.Errorf("puzzle name doesn't match. got: %v expected: %v", secondPuzzle.Name, secondName)
	}

	// Test "print"; should match input.
	output := bytes.NewBuffer(make([]byte, 0, len(secondCase)))
	secondPuzzle.Print(output)

	// NB: Print always terminates with a newline, but it doesn't care whether there's a trailing newline.
	if output.String() != (secondCase + "\n") {
		t.Errorf("Print failed:\ngot:\n%v\nexpected:\n%v\n---\n", output, secondCase)
	}
}
