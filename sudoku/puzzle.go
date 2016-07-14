package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
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
		if err == io.EOF {
			return nil, fmt.Errorf("reached EOF early after reading %q: %v", name, err)
		}
		return nil, err
	}
	name = strings.Trim(name, "\n")

	var grid []uint64
	dimSq := dimension * dimension
	for x := 0; x < dimSq; x++ {
		// Read a line into the grid, parsing it to ints.
		line, err := buf.ReadString('\n')
		if err != nil && err != io.EOF {
			// EOF is okay; we read that at the last line.
			return nil, err
		}
		line = strings.Trim(line, "\n")

		if len(line) != dimSq {
			return nil, fmt.Errorf("Line %q for case %q has %d elements, not %d", line, name, len(line), dimSq)
		}

		var c rune
		for _, c = range line {
			// https://blog.golang.org/strings :
			// "A for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration."
			// And yet, the default formatter for rune appears to be the codepoint's number. Huh?
			// Print the rune back out as a string; and then parse it as a Uint in base-36.
			s := fmt.Sprintf("%c", c)
			i, err := strconv.ParseUint(s, base, 64)

			if err != nil {
				return nil, err
			}

			if i > uint64(dimSq) {
				return nil, fmt.Errorf("Value %d (%q) is outside of the limits of dimension %d", i, s, dimension)
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
	dimSq := p.dimension * p.dimension
	for i := 0; i < dimSq; i++ {
		// Columns
		for j := 0; j < dimSq; j++ {
			v := p.grid[i * dimSq + j]
			s := strconv.FormatUint(v, base)
			fmt.Fprint(w, s)
		}
		fmt.Fprintln(w, "")
	}
}
