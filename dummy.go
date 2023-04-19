package lazyreader

import (
	"io"
)

type (
	EmptyReader struct {
		Length int
		offset int
	}
)

func (x *EmptyReader) Read(b []byte) (n int, err error) {
	l := len(b)
	if x.offset == x.Length {
		return 0, io.EOF
	}
	if x.offset+l > x.Length {
		d := x.Length - x.offset
		x.offset += d
		return d, nil
	}
	x.offset += l
	return l, nil
}
