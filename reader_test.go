package lazyreader

import (
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	length := 1024*8 + 123 // about 8KB
	empty := &EmptyReader{Length: length}
	x := NewWithConfig(empty, 4*8*1000, 0) // 4KB/s -> 32kbps
	n, err := io.Copy(io.Discard, x)       // will take about 2s (8/4)
	if err != nil {
		t.Fatal(err)
	}
	if n != int64(length) {
		t.Error("unexpected read length:", n, "expected:", length)
	}
}
