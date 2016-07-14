package sudoku

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestPuzzleCollection(t *testing.T) {
	prefixReader := strings.NewReader(tradPrefix)
	firstReader := strings.NewReader(firstCase)
	secondReader := strings.NewReader(secondCase)
	catReader := io.MultiReader(prefixReader, firstReader, secondReader)

	expectedOutput := new(bytes.Buffer)
	r := io.TeeReader(catReader, expectedOutput)

	collection, err := NewCollection(r)
	if err != nil {
		t.Fatalf("couldn't create PuzzleCollection: %v", err)
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
