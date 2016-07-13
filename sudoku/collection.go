package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const(
	maxDim = 5 // 1-9A-Z allows for at most 35 characters, so at most a 5x5 puzzle. 
)

type PuzzleCollection struct {
	puzzles map[string]*Puzzle
	dimension int
}

// Read per the input format.
func NewCollection(r io.Reader) (*PuzzleCollection, error) {
	buf := bufio.NewReader(r)

	// Scan for two ints: the dimension, and the number of test cases.
	dimStr, err := buf.ReadString(' ')
	if err != nil {
		return nil, err
	}

	dimension, err := strconv.Atoi(dimStr)
	if err != nil {
		return nil, err
	}


	if dimension > maxDim {
		return nil, fmt.Errorf("Dimension %d of input exceeds maximum dimension %d", dimension, maxDim)
	}

	countStr, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}


	results := make(map[string]*Puzzle)

	for i := 0; i < count; i++ {
		puzzle, err := NewPuzzle(dimension, r)
		if err != nil {
			return nil, err
		}

		if _, ok := results[puzzle.name]; ok {
			return nil, fmt.Errorf("Duplicate puzzle name %s", puzzle.name)
		}
		results[puzzle.name] = puzzle
	}

	return &PuzzleCollection{
		puzzles: results,
		dimension: dimension,
	}, nil
}

func (p *PuzzleCollection) Print(w io.Writer) {
	fmt.Printf("%d %d\n", p.dimension, len(p.puzzles))
	for _, puzzle := range p.puzzles {
		puzzle.Print(w)
	}
}
