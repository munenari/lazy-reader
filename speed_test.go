package lazyreader

import "testing"

func TestCalcDuration(t *testing.T) {
	if res := calcDuration(10, 10); res.String() != "1s" {
		t.Error(res)
	}
	if res := calcDuration(10, 1); res.String() != "10s" {
		t.Error(res)
	}
	if res := calcDuration(1, 10); res.String() != "100ms" {
		t.Error(res)
	}
	if res := calcDuration(1, 0); res.String() != "0s" {
		t.Error(res)
	}
}
