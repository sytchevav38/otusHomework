package main

import "fmt"

func main() {
	a := 9

	someFunc(a)
	//someFuncPtr(&a)
	fmt.Println(a)
	a = someFuncV2(a)
	fmt.Println(a)
}

func someFunc(a int) {
	a += 100
	fmt.Println(a)
}

func someFuncPtr(a *int) {
	if a != nil {
		*a = 100
	}
}

func someFuncV2(a int) int {
	a += 100
	return a
}
