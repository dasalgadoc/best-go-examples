package main

import (
	"fmt"
)

type WorkerPool struct {
	tasks   []int
	workers int
	jobs    chan int // IO Channel
	results chan int // IO Channel
}

func NewWorkerPool(task []int, workers int) WorkerPool {
	return WorkerPool{
		tasks:   task,
		workers: workers,
		jobs:    make(chan int, len(task)),
		results: make(chan int, len(task)),
	}
}

func (w *WorkerPool) Run() {
	// Concurrent threads to read shared channels
	for i := 0; i < w.workers; i++ {
		go worker(i, w.jobs, w.results)
	}

	// Fill with task
	for _, task := range w.tasks {
		w.jobs <- task
	}
	close(w.jobs)

	for i := 0; i < len(w.tasks); i++ {
		<-w.results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("--> Worker #%d started fibonacci with %d\n", id, job)
		fib := fibonacci(job)
		fmt.Printf("Worker #%d finished. Job: %d; fibonacci: %d\n", id, job, fib)
		results <- fib
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
