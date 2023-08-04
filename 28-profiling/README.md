# 🪚 Go Profiling

## 📔 Method 1: pprof with commands

1. Run test 

```bash
go test -cpuprofile=cpu.prof -memprofile=mem.prof 
```

2. Inspect the generated files

```bash
go tool pprof cpu.prof
```

or

```bash
go tool pprof mem.prof
```

This commands will open an interface in the terminal. Some functions are:

|      Command      | Description                                   |
|:-----------------:|:----------------------------------------------|
|       `top`       | Show the top entries                          |
| `list <function>` | Show the code of the function                 |
|       `web`       | Open a web interface to inspect the profile * |
|       `pdf`       | Generate a pdf file                           |
|      `help`       | Show the help                                 |
|      `quit`       | Exit the interface                            |

* _Note_: This command require [Graphviz](https://www.graphviz.org/) installed in the system. 

## 📚Method 2: pprof embedded in the code

For cpu profiling write at the beginning of the function to profile:

```go
func functionToProfile() {
    // Write CPU profile to file
    cpu, _ := os.Create("cpu.prof")
    defer cpu.Close()
    pprof.StartCPUProfile(cpu)
    defer pprof.StopCPUProfile()

    // Do something
}
```

For memory profiling write at the end of the function to profile:

```go
func functionToProfile() {
    // Do something

    // Write memory profile to file
    mem, _ := os.Create("mem.prof")
    defer mem.Close()
    pprof.WriteHeapProfile(mem)
}
```
