package main

import "fmt"

type figure interface {
	area() float64
}

func calculateArea(f figure) {
	fmt.Printf("Area of %T: %f\n", f, f.area())
}

type square struct {
	base float64
}

func (s square) area() float64 {
	return s.base * s.base
}

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.height * r.base
}

// CypherAlgorithm composition
type CypherAlgorithm interface {
	Encrypt(string) string
	Decrypt(string) string
}

type AES struct {
	CypherAlgorithm
}

func (a *AES) Encrypt(entry string) string {
	return "AES Encrypt: " + entry
}

func (a *AES) Decrypt(entry string) string {
	return "AES Decrypt: " + entry
}

type RSA struct {
	algo CypherAlgorithm
}

func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	calculateArea(mySquare)
	calculateArea(myRectangle)

	// a interface can be any struct or variable
	myInterface := []interface{}{"Hello", 12, 4.98, true}
	fmt.Println(myInterface...)

	// composition
	aes := AES{}
	fmt.Println(aes.Encrypt("Hello"))
	fmt.Println(aes.Decrypt("Hello"))

	rsa := RSA{algo: &aes}
	fmt.Println(rsa.algo.Encrypt("Hello"))
	fmt.Println(rsa.algo.Decrypt("Hello"))
}
