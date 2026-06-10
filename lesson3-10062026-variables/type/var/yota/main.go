package main

import "fmt"

const (
	A = iota + 25
	B
	C
	D
)

func main() {
	fmt.Println(A, B, C, D)
}
