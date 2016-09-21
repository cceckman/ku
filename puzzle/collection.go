package puzzle

import (
	"bufio"
	"bytes"
	"strings"
	"strconv"
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
	buf := bytes.NewBufferString(s.Text())
	sp := bufio.NewScanner(buf)
	sp.Split(bufio.ScanWords)
	// Get name
	if ok := sp.Scan(); ! ok {
		return int64(cr.Count), sp.Err()
	}
	c.Name = strings.TrimSpace(sp.Text())
	// Get count
	if ok := sp.Scan(); ! ok {
		return int64(cr.Count), sp.Err()
	}
	sz, err := strconv.ParseInt(sp.Text(), 10, 32)
	if err != nil {
		return int64(cr.Count), fmt.Errorf("error in parsing size '%s' of collection '%s': %v",
			sp.Text(), c.Name, err)
	}
	c.Puzzles = make([]*Puzzle, sz)

	// While there are lines to scan
	for i := range c.Puzzles {
		p, err := NewPuzzle(cr)
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
	cw := NewCountingWriter(w)

	_, err := fmt.Fprintf(w, "%s %d\n", p.Name, len(p.Puzzles))
	if err != nil {
		return int64(cw.Count), err
	}

	for _, puzzle := range p.Puzzles {
		_, err := puzzle.WriteTo(w)
		if err != nil {
			return int64(cw.Count), err
		}
	}

	return int64(cw.Count), nil
}
