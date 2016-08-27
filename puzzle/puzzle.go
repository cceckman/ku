package puzzle

import (
	"fmt"
	"io"
	"strings"
)

type Puzzle struct {
	Name string
	// The size, e.g. 3 for a standard Sudoku. Only "square" Sudoku are supported.
	Size int
	// Size * Size
	len int
	// TODO : Support >2 dimensions.
	Value []uint64
}

// Interface assertions. This appears to be the Go-ish way to assert "I implement an interface."
var (
	_ io.ReaderFrom = (*Puzzle)(nil)
	_ io.WriterTo   = (*Puzzle)(nil)
)

// Load a single puzzle from the Reader.
// The format is:
// The name, on a line on its own
// The size (e.g. 3) on a line on its own.
// The data, with optional newlines within it, terminating with a newline.
func NewPuzzle(r io.Reader) (*Puzzle, error) {
	p := &Puzzle{}
	if _, err := p.ReadFrom(r); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Puzzle) ReadFrom(r io.Reader) (int64, error) {
	name := new(string)
	if c, err := fmt.Fscanln(r, "%s", name); c != 1 || err != nil {
		return int64(len(*name)), fmt.Errorf("error in scanning collection name: read %d, error: %v", c, err)
	}
	p.Name = strings.TrimSpace(*name)

	if c, err := fmt.Fscanf(r, "%d", &p.Size); c != 1 || err != nil {
		return int64(c), fmt.Errorf("error in scanning size of puzzle '%s': read %d values, error: %v", p.Name, c, err)
	}

	if p.Size < 1 || p.Size > 3 {
		return 0, fmt.Errorf("Size %d of puzzle '%s' is not supported", p.Size, p.Name)
	}
	p.len = p.Size * p.Size

	p.Value = make([]uint64, p.len)
	acc := int64(0)

	for x := 0; x < p.len; x++ {
		c, err := fmt.Fscanf(r, "%x", &p.Value[x])
		acc += int64(c)
		if c != 1 || err != nil {
			return acc, fmt.Errorf("error in reading position %d of puzzle '%s': read %d values, error: %v", x, name, c, err)
		}
	}

	return acc, nil
}

// Prints a puzzle to the Writer.
func (p *Puzzle) WriteTo(w io.Writer) (int64, error) {
	var acc int64
	// Name
	sz, err := fmt.Fprintf(w, "%s\n", p.Name)
	acc += int64(sz)
	if err != nil {
		return acc, err
	}
	// Size
	sz, err = fmt.Fprintf(w, "%d\n", p.Size)
	acc += int64(sz)
	if err != nil {
		return acc, err
	}

	// Lines
	for i := 0; i < p.len; i++ {
		// Columns
		for j := 0; j < p.len; j++ {
			v := p.Value[i*p.len+j]
			sz, err = fmt.Fprintf(w, "%x ", v)
			acc += int64(sz)
			if err != nil {
				return acc, err
			}
		}
		sz, err := fmt.Fprintln(w, "")
		acc += int64(sz)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// RowOf gives the row of the given cell.
func (p *Puzzle) RowOf(cell int) int {
	return cell / p.len
}

// Row gives the indeces of the cells in the given row.
func (p *Puzzle) Row(row int) []int {
	r := make([]int, p.len)
	for i := 0; i < p.len; i++ {
		r[i] = row*p.len + i
	}
	return r
}

// ColOf gives the column of the given cell.
func (p *Puzzle) ColOf(cell int) int {
	return cell % p.len
}

// Col gives the indeces of the cells in the given row.
func (p *Puzzle) Col(col int) []int {
	r := make([]int, p.len)
	for i := 0; i < p.len; i++ {
		r[i] = i*p.len + col
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

	r := make([]int, p.len)

	// Base row for this box
	//bRow := (box / p.Size) * p.Size
	bRow := box - (box % p.Size)
	// Base column for this box
	bCol := (box % p.Size) * p.Size
	for i := 0; i < p.len; i++ {
		col := bCol + (i % p.Size)
		row := bRow + (i / p.Size)
		r[i] = row*p.len + col
	}
	return r
}

// Get all the values at the given indices.
func (p *Puzzle) Values(idx []int) []uint64 {
	results := make([]uint64, len(idx))
	for i, n := range idx {
		results[i] = p.Value[n]
	}
	return results
}

// Create a new copy of the puzzle.
func (p *Puzzle) Copy() *Puzzle {
	newp := &Puzzle{
		Name:  p.Name,
		Size:  p.Size,
		Value: make([]uint64, len(p.Value)),
	}
	for i := range p.Value {
		newp.Value[i] = p.Value[i]
	}
	return newp
}

// Utility: pretty-print an index, for debugging.
func (p *Puzzle) CellInfo(idx int) string {
	return fmt.Sprintf("idx: %d value: %d r: %d c: %d b: %d",
		idx, p.Value[idx], p.RowOf(idx), p.ColOf(idx), p.BoxOf(idx))
}
