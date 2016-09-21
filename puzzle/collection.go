package puzzle

import (
	"bufio"
	"fmt"
	"io"
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

// Read a Collection. 
// The first line contains the name of the collection (which has no spaces),
// followed by a space, followed by the number of puzzles in this collection.
// Then, the rest of the lines describe Puzzles per the Puzzle format.
func NewCollection(r io.Reader) (*Collection, error) {
	c := &Collection{}
	if _, err := c.ReadFrom(r); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Collection) ReadFrom(r io.Reader) (int64, error) {
	cr := NewCountingReader(r)
	s := bufio.NewScanner(cr)
	s.Split(bufio.ScanLines)

	if ok := s.Scan(); ! ok {
		return int64(cr.Count), s.Err()
	}

	// Split the first line by space...
	_ = s.Text()
	// DO NOT SUBMIT - TODO

	// While there are lines to scan
	for i := 0; i < len(c.Puzzles); i++ {
		p, err := NewPuzzle(r)
		if err != nil {
			return int64(cr.Count), fmt.Errorf("error reading puzzle %d in collection '%s': %v",
				i, c.Name, err)
		}
		c.Puzzles[i] = p
	}

	return int64(cr.Count), nil
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
