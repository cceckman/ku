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

type CountingReader struct {
	wrap io.Reader
	Count int
}

func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader { wrap: r }
}

func (c *CountingReader) Read(p []byte) (n int, err error) {
	n, err = c.wrap.Read(p)
	c.Count += n
	return
}

type CountingWriter struct {
	wrap io.Writer
	Count int
}

func NewCountingWriter(r io.Writer) *CountingWriter {
	return &CountingWriter { wrap: r }
}

func (c *CountingWriter) Write(p []byte) (n int, err error) {
	n, err = c.wrap.Write(p)
	c.Count += n
	return
}
