package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Puzzle struct {
	name      string
	dimension int
	grid      []uint64
}

// Load a single puzzle from the Reader.
func NewPuzzle(dimension int, r io.Reader) (*Puzzle, error) {
	buf := bufio.NewReader(r)
	name, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var grid []uint64
	for x := 0; x < dimension; x++ {
		// Read a line into the grid, parsing it to ints.
		line, err := buf.ReadString('\n')
		if err != nil {
			return nil, err
		}

		if len(line) != dimension {
			return nil, fmt.Errorf("Line for case %q has %d elements, not %d", name, len(line), dimension)
		}

		for _, c := range line {
			i, err := strconv.ParseUint(fmt.Sprint(c), 36, 64)

			if err != nil {
				return nil, err
			}

			n := uint64(dimension)
			if i > n * n {
				return nil, fmt.Errorf("Value %d is outside of the limits of dimension %d", i, dimension)
			}

			grid = append(grid, i)
		}
	}

	return &Puzzle{
		name:      name,
		dimension: dimension,
		grid:      grid,
	}, nil
}

// Prints a puzzle to the Writer.
func (p *Puzzle) Print(w io.Writer) {
	fmt.Fprintf(w, "%s\n", p.name)
	// Lines
	for i := 0; i < p.dimension; i++ {
		// Columns
		for j := 0; j < p.dimension; j++ {
			fmt.Fprintf(w, "%c", p.grid[i*p.dimension+j])
		}
		fmt.Fprintln(w, "")
	}
}
