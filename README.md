# Mist 薄雾算法
不同于 snowflake 的全局唯一 ID 生成算法及其 Golang 实现。


### Benchmark 测试结果

进行了多轮测试，随机取 3 轮测试结果。以此计算平均值，得 `单次执行时间 8981 ns`。以下是随机 3 轮测试的结果：


```
goos: darwin
goarch: amd64
pkg: mist
BenchmarkGenerate-4       118484              8863 ns/op
PASS
ok      mist    1.345s
```

```
goos: darwin
goarch: amd64
pkg: mist
BenchmarkGenerate-4       132276              9038 ns/op
PASS
ok      mist    1.539s
```

```
goos: darwin
goarch: amd64
pkg: mist
BenchmarkGenerate-4       116197              9042 ns/op
PASS
ok      mist    1.388s
```


