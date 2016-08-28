package puzzle

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

const (
	collectionCase = `Test Collection
2
Case 1
3
0 0 0 0 0 0 0 0 1 
6 0 3 0 9 1 0 0 5 
0 7 9 0 4 0 0 8 0 
0 5 0 0 7 4 0 0 0 
0 0 0 0 0 2 0 0 6 
0 0 0 0 3 0 0 0 0 
5 0 4 0 6 0 0 9 0 
0 0 6 0 0 8 0 0 4 
3 0 0 0 0 0 7 0 0
Case 4
3
1 9 8 7 3 4 2 6 5 
5 6 4 1 9 2 3 7 8 
2 7 3 8 6 5 9 1 4 
3 1 5 4 2 7 6 8 9 
8 4 9 3 5 6 7 2 1 
6 2 7 9 1 8 5 4 3 
7 3 6 5 4 1 8 9 2 
4 5 2 6 8 9 1 3 7 
9 8 1 2 7 3 4 5 6
`
	firstName  = "Case 1"
	secondName = "Case 4"
	// TODO add >3x3 case
	firstCase = `Case 1
3
0 0 0 0 0 0 0 0 1 
6 0 3 0 9 1 0 0 5 
0 7 9 0 4 0 0 8 0 
0 5 0 0 7 4 0 0 0 
0 0 0 0 0 2 0 0 6 
0 0 0 0 3 0 0 0 0 
5 0 4 0 6 0 0 9 0 
0 0 6 0 0 8 0 0 4 
3 0 0 0 0 0 7 0 0
`
)

func TestCollection(t *testing.T) {
	// TODO test invalid collections, e.g. 2 of the same name
	collectionReader := strings.NewReader(collectionCase)

	expectedOutput := new(bytes.Buffer)
	r := io.TeeReader(collectionReader, expectedOutput)

	collection, err := NewCollection(r)
	if err != nil {
		t.Fatalf("couldn't create Collection: %v", err)
	}

	// Test "print"; should match the input read.
	output := new(bytes.Buffer)
	collection.WriteTo(output)

	// NB: Print always terminates with a newline, but it doesn't care whether there's a trailing newline.
	expectedOutput.WriteString("\n")
	if output.String() != expectedOutput.String() {
		t.Errorf("Print failed:\ngot:\n%v\nexpected:\n%v\n---\n", output, expectedOutput.String())
	}
}

func TestSinglePuzzle(t *testing.T) {
	firstReader := strings.NewReader(firstCase)
	p, err := NewPuzzle(firstReader)
	if err != nil {
		t.Errorf("got error when loading first puzzle: %v", err)
	}

	// Check properties
	if p.Name != firstName {
		t.Errorf("puzzle name doesn't match. got: %v expected: %v", p.Name, firstName)
	}
	if p.Size != 3 {
		t.Errorf("size doesn't match. got: %v expected: %v", p.Size, 3)
	}
	totallen := (3 * 3) * (3 * 3) // Two-dimensional, with size 3)
	if len(p.Value) != totallen {
		t.Errorf("data dimension doesn't match. got: %v expected: %v", len(p.Value), totallen)
	}

	dump := false

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
	// Spot-check Row:
	for idx, row := range map[int][]uint64{
		0: []uint64{0, 0, 0, 0, 0, 0, 0, 0, 1},
		6: []uint64{5, 0, 4, 0, 6, 0, 0, 9, 0},
		8: []uint64{3, 0, 0, 0, 0, 0, 7, 0, 0},
	} {
		gotRow := p.Row(idx)
		if len(gotRow) != len(row) {
			t.Errorf("Row %d test failed: got: %v expected: %v", idx, gotRow, row)
		}
		for i := range row {
			if p.Value[gotRow[i]] != row[i] {
				t.Errorf("Row %d test failed: got: (%v)  expected value: %v", idx, p.CellInfo(gotRow[i]), row[i])
				dump = true
			}
		}
	}
	// Spot-check Col:
	for idx, col := range map[int][]uint64{
		0: []uint64{0, 6, 0, 0, 0, 0, 5, 0, 3},
		4: []uint64{0, 9, 4, 7, 0, 3, 6, 0, 0},
	} {
		gotCol := p.Col(idx)
		if len(gotCol) != len(col) {
			t.Errorf("Col %d test failed: got: %v expected: %v", idx, gotCol, col)
		}
		for i := range col {
			if p.Value[gotCol[i]] != col[i] {
				t.Errorf("Col %d test failed: got: (%v)  expected value: %v", idx, p.CellInfo(gotCol[i]), col[i])
				dump = true
			}
		}
	}
	// Spot-check Box:
	for idx, box := range map[int][]uint64{
		0: []uint64{0, 0, 0, 6, 0, 3, 0, 7, 9},
		3: []uint64{0, 5, 0, 0, 0, 0, 0, 0, 0},
		8: []uint64{0, 9, 0, 0, 0, 4, 7, 0, 0},
	} {
		gotBox := p.Box(idx)
		if len(gotBox) != len(box) {
			t.Errorf("Box %d test failed: got: %v expected: %v", idx, gotBox, box)
		}
		for i := range box {
			if p.Value[gotBox[i]] != box[i] {
				t.Errorf("Box %d test failed: got: (%v)  expected value: %v", idx, p.CellInfo(gotBox[i]), box[i])
				dump = true
			}
		}
	}

	// Test "WriteTo"; should match input.
	output := bytes.NewBuffer(make([]byte, 0, len(firstCase)))
	p.WriteTo(output)

	// NB: Print always terminates with a newline, but it doesn't care whether there's a trailing newline.
	if output.String() != firstCase {
		t.Errorf("Print failed:\ngot:\n%v\nexpected:\n%v\n---\n", output, firstCase)
	}

	// Write anyway if 'dump' was triggered
	if dump {
		t.Logf("Printed value: \n%v", output)
	}

}
