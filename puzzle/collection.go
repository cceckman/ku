package puzzle

import (
	"fmt"
	"io"
	"strings"
)

type Collection struct {
	Name    string
	Puzzles []*Puzzle
}

// Interface assertions. This appears to be the Go-ish way to assert "I implement an interface."
var (
	_ io.ReaderFrom = (*Collection)(nil)
	_ io.WriterTo   = (*Collection)(nil)
)

// Read a Collection. The input format is as follows:
// The first line gives the name of the Collection.
// The second line has one number: the number of puzzles in this Collection.
// Then, the rest of the lines describe Puzzles per the Puzzle format.
func NewCollection(r io.Reader) (*Collection, error) {
	c := &Collection{}
	if _, err := c.ReadFrom(r); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Collection) ReadFrom(r io.Reader) (int64, error) {
	name := new(string)
	if count, err := fmt.Fscanln(r, "%s", name); count != 1 || err != nil {
		return int64(len(*name)), fmt.Errorf("error in scanning collection name: read %d, error: %v", count, err)
	}
	c.Name = strings.TrimSpace(*name)

	var size int
	if count, err := fmt.Fscanf(r, "%d", &size); count != 1 || err != nil {
		return int64(count), fmt.Errorf("error in scanning size of collection '%s': read %d values, error: %v", c.Name, count, err)
	}

	c.Puzzles = make([]*Puzzle, size)

	acc := int64(0)
	for i := 0; i < size; i++ {
		n, err := c.Puzzles[i].ReadFrom(r)
		acc += n
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Makes a copy of this PuzzleCollection.
func (c *Collection) Copy() *Collection {
	r := &Collection{
		Puzzles: make([]*Puzzle, len(c.Puzzles)),
		Name:    c.Name,
	}
	for i, p := range c.Puzzles {
		r.Puzzles[i] = p.Copy()
	}
	return r
}

func (p *Collection) WriteTo(w io.Writer) (int64, error) {
	var acc int64
	sz, err := fmt.Fprintf(w, "%s\n%d\n", p.Name, len(p.Puzzles))
	acc += int64(sz)
	if err != nil {
		return acc, err
	}

	for _, puzzle := range p.Puzzles {
		sz, err := puzzle.WriteTo(w)
		acc += int64(sz)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}
