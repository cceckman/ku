package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Puzzle struct {
	name string
	Size int
	grid []uint64 // a Size-by-Size grid in row-major order.
}

// Load a single puzzle from the Reader.
func NewPuzzle(size int, r io.Reader) (*Puzzle, error) {
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
	dimSq := size * size
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
				return nil, fmt.Errorf("Value %d (%q) is outside of the limits of size %d", i, s, size)
			}

			grid = append(grid, i)
		}
	}

	return &Puzzle{
		name: name,
		Size: size,
		grid: grid,
	}, nil
}

// Prints a puzzle to the Writer.
func (p *Puzzle) Print(w io.Writer) {
	fmt.Fprintf(w, "%s\n", p.name)
	// Lines
	dimSq := p.Size * p.Size
	for i := 0; i < dimSq; i++ {
		// Columns
		for j := 0; j < dimSq; j++ {
			v := p.grid[i*dimSq+j]
			s := strconv.FormatUint(v, base)
			fmt.Fprint(w, s)
		}
		fmt.Fprintln(w, "")
	}
}

// RowOf gives the row of the given cell.
func (p *Puzzle) RowOf(cell int) int {
	return cell / (p.Size * p.Size)
}

// Row gives the indeces of the cells in the given row.
func (p *Puzzle) Row(row int) []int {
	dimSq := p.Size * p.Size
	r := make([]int, dimSq)
	for i := 0; i < dimSq; i++ {
		r[i] = row*dimSq + i
	}
	return r
}

// ColOf gives the column of the given cell.
func (p *Puzzle) ColOf(cell int) int {
	return cell % (p.Size * p.Size)
}

// Col gives the indeces of the cells in the given row.
func (p *Puzzle) Col(col int) []int {
	dimSq := p.Size * p.Size
	r := make([]int, p.Size)
	for i := 0; i < dimSq; i++ {
		r[i] = i*dimSq + col
	}
	return r
}

// BoxOf gives the "box" of the given cell- which NxN subdivision the cell is in.
// Subdivisions are numbered from the top-left, starting at zero, row-major order.
func (p *Puzzle) BoxOf(cell int) int {
	// Grid, with Row and Col:
	// 	  0  1  2  3  4  5  6  7  8
	//0	  0  1  2  3  4  5  6  7  8
	//1	  9 10 11 12 13 14 15 16 17
	//2	 18 19 20 21 22 23 24 25 26
	//3	 27 28 29 30 31 32 33 34 35
	// Col / size, row / size
	// 	  0  0  0  1  1  1  2  2  2
	//0	  0  1  2  3  4  5  6  7  8
	//0	  9 10 11 12 13 14 15 16 17
	//0	 18 19 20 21 22 23 24 25 26
	//1	 27 28 29 30 31 32 33 34 35
	// Col / size + (row / size) * size
	// 	  0  0  0  1  1  1  2  2  2
	//0	  0  0  0  1  1  1  6  7  8
	//0	  0 00  0	 1  1  1 15 16 17
	//0	 00 00 00  1  1  1 24 25 26
	//3	  3 28 29  4 31 32 33 34 35
	// row - (row % size) is easier to compute- mod and sub are faster operations.
	row := p.RowOf(cell)
	col := p.ColOf(cell)
	return (col / p.Size) + (row - (row % p.Size))
}

// Box gives the indices of cells in the given box.
// Boxes are indexed in English-order: left to right, then the next row. 0 is in the upper-left.
func (p *Puzzle) Box(box int) []int {
	// 	  0  0  0  1  1  1  2  2  2
	//0	  0  1  2  3  4  5  6  7  8
	//0	  9 10 11 12 13 14 15 16 17
	//0	 18 19 20 21 22 23 24 25 26
	//1	 27 28 29 30 31 32 33 34 35

	dimSq := p.Size * p.Size
	r := make([]int, p.Size)

	// Base row for this box
	//bRow := (box / p.Size) * p.Size
	bRow := box - (box % p.Size)
	// Base column for this box
	bCol := (box % p.Size) * p.Size
	for i := 0; i < dimSq; i++ {
		col := bCol + (i % p.Size)
		row := bRow + (i / p.Size)
		r[i] = row*dimSq + col
	}
	return r
}
