package gopl

import (
	"io"
)

// StringReader implement io.Reader interface
// and wrap a source string as a Reader.
// The Strings.Reader also contain a 'prevRune' field.
type StringReader struct {
	s   string
	idx int
}

func NewReader(source string) *StringReader {
	return &StringReader{s: source}
}

func (sr *StringReader) Read(p []byte) (int, error) {
	if sr.idx >= len(sr.s) {
		return 0, io.EOF
	}
	n := copy(p, sr.s[sr.idx:])
	sr.idx += n
	return n, nil
}
