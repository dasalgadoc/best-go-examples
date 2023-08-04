# 🔎 Benchmarking

To benchmark and compare the performance of the different implementations.
It's necessary to implement a `benchmark` from the package testing.

A benchmark is a function that calls the code to be measured several times. The function must start with the word `Benchmark` and take one argument of type *testing.B.

```go
func BenchmarkSomething(b *testing.B) {
	b.ResetTimer()
	// N is the number of times the code inside the loop is executed.
    for i := 0; i < b.N; i++ {
        // Your code here
    }
}
```

## 📌 Get the results

To get the results of the benchmark, you need to run the following command:

```bash
go test -bench=. 
```

or

```bash
go test -bench=. -benchmem
```
