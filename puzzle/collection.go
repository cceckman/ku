package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	maxdim = 5 // 1-9A-Z allows for at most 35 characters, so at most a 5x5 puzzle.
	base   = maxdim * maxdim
)

type Collection struct {
	Puzzles []*Puzzle
	Size    int
}

// Interface assertions. This appears to be the Go-ish way to assert "I implement an interface."
var (
	_ io.WriterTo = (*Collection)(nil)
)

// Read per the input format.
func NewCollection(r io.Reader) (*Collection, error) {
	buf := bufio.NewReader(r)

	// Scan for two ints: the size, and the number of test cases.
	dimStr, err := buf.ReadString(' ')
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(strings.Trim(dimStr, " "))
	if err != nil {
		return nil, err
	}

	if size > maxdim {
		return nil, fmt.Errorf("dimension %d of input exceeds maximum size %d", size, maxdim)
	}

	countStr, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}

	count, err := strconv.Atoi(strings.Trim(countStr, "\n "))
	if err != nil {
		return nil, err
	}

	result := &Collection{Size: size}
	names := make(map[string]bool)
	for i := 0; i < count; i++ {
		// Pass in the already-buffered reader, as it may have already forwarded the read pointer in 'r'
		// past where we're interestd in.
		puzzle, err := NewPuzzle(size, buf)
		if err != nil {
			return nil, err
		}

		if _, ok := names[puzzle.Name]; ok {
			return nil, fmt.Errorf("Duplicate puzzle name %s", puzzle.Name)
		}
		result.Puzzles = append(result.Puzzles, puzzle)
		names[puzzle.Name] = true
	}

	return result, nil
}

// Makes a copy of this PuzzleCollection.
func (c *Collection) Copy() (*Collection) {
	r := &Collection{
		Puzzles: make([]*Puzzle, len(c.Puzzles)),
		Size: c.Size,
	}
	for i, p := range c.Puzzles {
		r.Puzzles[i] = p.Copy()
	}
	return r
}

func (p *Collection) WriteTo(w io.Writer) (int64, error) {
	var acc int64
	sz, err := fmt.Fprintf(w, "%d %d\n", p.Size, len(p.Puzzles))
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
