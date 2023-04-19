package lazyreader

import "testing"

func TestEmptyReader(t *testing.T) {
	b := make([]byte, 8)
	t.Run("le", func(t *testing.T) {
		x := &EmptyReader{Length: 10}
		if n, _ := x.Read(b); n != 8 {
			t.Error("unexpected length", n)
		}
	})
	t.Run("gt", func(t *testing.T) {
		x := &EmptyReader{Length: 10}
		if n, _ := x.Read(b); n != 8 {
			t.Error("unexpected length", n)
		}
		if n, _ := x.Read(b); n != 2 {
			t.Error("unexpected length", n)
		}
	})
	t.Run("gt", func(t *testing.T) {
		x := &EmptyReader{Length: 4}
		if n, _ := x.Read(b); n != 4 {
			t.Error("unexpected length", n)
		}
	})
	t.Run("eq", func(t *testing.T) {
		x := &EmptyReader{Length: 8}
		if n, _ := x.Read(b); n != 8 {
			t.Error("unexpected length", n)
		}
	})
}
