package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Write CPU profile to file
	cpu, _ := os.Create("./28-profiling/embedded.cpu.prof")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	fmt.Println("Starting...")
	start := time.Now()
	Fibonacci(50)
	fmt.Println("Finished")
	elapsed := time.Since(start)
	fmt.Printf("Ellapsed time: %s\n", elapsed)

	// Write memory profile to file
	mem, _ := os.Create("./28-profiling/embedded.mem.prof")
	defer mem.Close()
	pprof.WriteHeapProfile(mem)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
