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
	puzzles []*Puzzle
	Size    int
}

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

		if _, ok := names[puzzle.name]; ok {
			return nil, fmt.Errorf("Duplicate puzzle name %s", puzzle.name)
		}
		result.puzzles = append(result.puzzles, puzzle)
		names[puzzle.name] = true
	}

	return result, nil
}

func (p *Collection) Print(w io.Writer) {
	fmt.Fprintf(w, "%d %d\n", p.Size, len(p.puzzles))
	for _, puzzle := range p.puzzles {
		puzzle.Print(w)
	}
}
