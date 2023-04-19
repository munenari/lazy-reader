package lazyreader

import (
	"time"
)

const (
	BPS1M   int = 1 * 1e6
	BPS10M  int = 10 * BPS1M
	BPS100M int = 10 * BPS10M
	BPS1G   int = 10 * BPS100M
	BPS10G  int = 10 * BPS1G
)

func calcDuration(bitsLength, bps int) time.Duration {
	if bps == 0 {
		return 0
	}
	return time.Duration((float64(bitsLength) / float64(bps) * float64(time.Second)))
}
