lazy-reader
====

reader with retricted byte speed

## example

```go
func main() {
	f, err := os.Open("/path/to/sample")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := lazyreader.NewWithConfig(f, lazyreader.BPS1M, 0)
	_, err = io.Copy(os.Discard, r) // will copy by 1Mbps
	if err != nil {
		panic(err)
	}
}
```
