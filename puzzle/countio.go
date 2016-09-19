//
// countio.go
// Copyright (C) 2016 cceckman <charles@cceckman.com>
//
// Wrap io.Reader and io.Writer with counters on read and write,
// to make it easier to implement ReadFrom and WriteTo.
//

package puzzle

import (
	"io"
)
var (
	_ io.Reader = (*CountingReader)(nil)
	_ io.Writer   = (*CountingWriter)(nil)
)

type CountingReader {
	wrap io.Reader
	Count int
}

func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader { wrap: r }
}

func (c *CountingReader) Read(p []byte) (n int, err error) {
	n, err = wrap.Read()
	c.Count += n
	return
}

type CountingWriteer {
	wrap io.Writeer
	Count int
}

func NewCountingWriteer(r io.Writeer) *CountingWriteer {
	return &CountingWriteer { wrap: r }
}

func (c *CountingWriteer) Write(p []byte) (n int, err error) {
	n, err = wrap.Write()
	c.Count += n
	return
}
