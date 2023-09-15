# Comparison of hashes

The list of hash functions for comparison:

* djb2
* sdbm
* sha1
* sha256
* murmur3 (github.com/spaolacci/murmur3)
* xxHash (github.com/cespare/xxhash)
* simhash (https://github.com/mfonda/simhash)


Command: `go test -bench=. -benchtime=5s`

## Header

```
goos: linux
goarch: amd64
pkg: compareHashes
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
```


## Results

Length of the text = 10K

```
BenchmarkSdbm-8          1000000              5146 ns/op
BenchmarkDjb2-8          1000000              5015 ns/op
BenchmarkSha256-8       31689594               190.8 ns/op
BenchmarkSha1-8         47700940               130.8 ns/op
BenchmarkMurmur3-8      962910477                6.082 ns/op
BenchmarkXXHash-8       968513018                6.237 ns/op
BenchmarkSimhash-8          5868           1005024 ns/op
```

Length of the text = 100K

```
BenchmarkSdbm-8           119226             50857 ns/op
BenchmarkDjb2-8           120360             49943 ns/op
BenchmarkSha256-8       32240737               190.2 ns/op
BenchmarkSha1-8         47900412               129.5 ns/op
BenchmarkMurmur3-8      978218433                6.005 ns/op
BenchmarkXXHash-8       942319635                6.346 ns/op
BenchmarkSimhash-8           445          13256837 ns/op
```

Length of the text = 1M

```
BenchmarkSdbm-8            10000            538815 ns/op
BenchmarkDjb2-8            10000            530542 ns/op
BenchmarkSha256-8       28745600               193.5 ns/op
BenchmarkSha1-8         46466902               137.2 ns/op
BenchmarkMurmur3-8      950074614                6.079 ns/op
BenchmarkXXHash-8       951294382                6.230 ns/op
BenchmarkSimhash-8            37         153814714 ns/op
```

Length of th text = 3M

```
BenchmarkSdbm-8             3955           1548929 ns/op
BenchmarkDjb2-8             3842           1494095 ns/op
BenchmarkSha256-8       30617457               196.2 ns/op
BenchmarkSha1-8         45319993               137.1 ns/op
BenchmarkMurmur3-8      982202215                6.155 ns/op
BenchmarkXXHash-8       925098828                6.451 ns/op
BenchmarkSimhash-8            12         479861198 ns/op
```