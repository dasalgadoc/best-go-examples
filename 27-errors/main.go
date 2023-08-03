package main

import (
	"errors"
	"fmt"
	"log"
)

var SentinelErrNotFound = errors.New("not found")

type TypeNotFoundError struct {
	s    string
	file string
	line int
}

func (e *TypeNotFoundError) Error() string {
	return fmt.Sprintf("%s: %s:%d", e.s, e.file, e.line)
}

func NewTypeNotFoundError(s string, file string, line int) error {
	return &TypeNotFoundError{s, file, line}
}

func main() {
	// Sentinel
	if _, err := fetchUserSentinel(1); err == SentinelErrNotFound {
		log.Fatal(err)
	}

	// Error type
	if _, err := fetchUserType(1); err != nil {
		if err, ok := err.(*TypeNotFoundError); ok {
			log.Fatal(err)
		}
	}
}

func fetchUserSentinel(id int) (string, error) {
	return "", SentinelErrNotFound
}

func fetchUserType(id int) (string, error) {
	return "", NewTypeNotFoundError("user not found", "main.go", 44)
}
