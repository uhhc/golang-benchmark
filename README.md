# golang-md5-benchmark

There are different methods to generate MD5 hash:

```go
func md5V1(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func md5V2(str string) string {
    hash := md5.Sum([]byte(str))
    md5str := hex.EncodeToString(hash[:])
    return md5str
}

func md5V3(str string) string {
    w := md5.New()
    _, _ = io.WriteString(w, str)
    md5str := fmt.Sprintf("%x", w.Sum(nil))
    return md5str
}

func md5V4(str string) string {
    data := []byte(str)
    md5str := fmt.Sprintf("%x", md5.Sum(data))
    return md5str
}
```

Here is the benchmark result:

```shell
# go test -bench=. -benchmem -run=none
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkMD5V1-12        5696600               206.6 ns/op            80 B/op          3 allocs/op
BenchmarkMD5V2-12        5831992               178.5 ns/op            64 B/op          2 allocs/op
BenchmarkMD5V3-12        3285093               365.3 ns/op           176 B/op          5 allocs/op
BenchmarkMD5V4-12        3197449               377.2 ns/op            64 B/op          3 allocs/op
```

Obviously, the best performance method is `md5V2()`.
