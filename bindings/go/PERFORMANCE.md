go test ./bindings/go -bench=.
goos: linux
goarch: amd64
pkg: github.com/supranational/blst/bindings/go
cpu: Intel(R) Core(TM) i7-4790K CPU @ 4.00GHz
BenchmarkVerifyFastAggregateUncompressedMinSig/1_G1-8               1114           1353581 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/10_G1-8              1108           1323160 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/50_G1-8               685           1540998 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/100_G1-8              634           1720144 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/300_G1-8              841           1361151 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/1000_G1-8            1124           1306114 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/4000_G1-8            1117           1410837 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/1_G2-8               1105           1508038 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/10_G2-8              1077           1318747 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/50_G2-8               931           1435618 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/100_G2-8              796           1372564 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/300_G2-8              806           1528979 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/1000_G2-8            1088           1444121 ns/op
BenchmarkVerifyFastAggregateUncompressedMinSig/4000_G2-8            1123           1265289 ns/op
BenchmarkCoreAggregateMinSig/1-8                                   13138             99396 ns/op
BenchmarkCoreAggregateMinSig/10-8                                   3064            355946 ns/op
BenchmarkCoreAggregateMinSig/50-8                                    938           1293479 ns/op
BenchmarkCoreAggregateMinSig/100-8                                   493           2432954 ns/op
BenchmarkCoreAggregateMinSig/300-8                                   170           7005713 ns/op
BenchmarkCoreAggregateMinSig/1000-8                                   50          23161016 ns/op
BenchmarkCoreAggregateMinSig/4000-8                                   12          92381103 ns/op
PASS
ok      github.com/supranational/blst/bindings/go       54.969s

