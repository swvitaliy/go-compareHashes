# Comparison of hashes

The list of hash functions for comparison:

* djb2
* sdbm
* sha1
* sha256
* murmur3 (github.com/spaolacci/murmur3)
* xxHash (github.com/cespare/xxhash)
* aesHash
* simhash (https://github.com/mfonda/simhash)


x86 processor supports of aes instructions:

```text
cat /proc/cpuinfo
...
flags		: fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault epb invpcid_single ssbd ibrs ibpb stibp ibrs_enhanced tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust sgx bmi1 avx2 smep bmi2 erms invpcid mpx rdseed adx smap clflushopt intel_pt xsaveopt xsavec xgetbv1 xsaves dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp md_clear flush_l1d arch_capabilities
...
```

Command:

```bash
go test -bench=. -benchtime=5s

```

## Results

Text length = 3M

```
goos: linux
goarch: amd64
pkg: compareHashes
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkSdbm-8             3742           1699309 ns/op
BenchmarkDjb2-8             3524           1673717 ns/op
BenchmarkSha256-8       26210958               220.8 ns/op
BenchmarkSha1-8         38899916               135.1 ns/op
BenchmarkMurmur3-8      926981096                6.151 ns/op
BenchmarkXXHash-8       934610797                6.770 ns/op
BenchmarkAes-8          478043392               11.83 ns/op
BenchmarkSimhash-8            12         491083351 ns/op
```