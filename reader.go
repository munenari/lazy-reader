package lazyreader

import (
	"io"
	"time"
)

type (
	LazyReader struct {
		Src           io.Reader
		BPS           int
		RemoteWaiting time.Duration

		pr      *io.PipeReader
		pw      *io.PipeWriter
		started bool
	}
)

func New(r io.Reader) *LazyReader {
	return NewWithConfig(r, BPS1G, 1*time.Millisecond)
}

func NewWithConfig(r io.Reader, bps int, waiting time.Duration) *LazyReader {
	x := &LazyReader{
		Src:           r,
		BPS:           bps,
		RemoteWaiting: waiting,
	}
	x.Start()
	return x
}

func (x *LazyReader) Start() {
	if x.started {
		return
	}
	pr, pw := io.Pipe()
	x.pr = pr
	x.pw = pw
	go func() {
		defer x.pw.Close()
		time.Sleep(x.RemoteWaiting)
		for {
			if x.pw == nil || x.Src == nil {
				break
			}
			t := time.Now()
			n, err := io.CopyN(x.pw, x.Src, 1024)
			nano := calcDuration(int(n)*8, x.BPS)
			time.Sleep(nano - time.Since(t))
			if err != nil {
				break
			}
		}
	}()
}

func (x *LazyReader) Read(b []byte) (n int, err error) {
	return x.pr.Read(b)
}

func (x *LazyReader) Close() error {
	x.pr.Close()
	x.pw.Close()
	x.started = false
	x.Src = nil
	return nil
}

var (
	_ io.ReadCloser = &LazyReader{}
)
