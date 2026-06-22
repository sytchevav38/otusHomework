package main

import "fmt"

func main() {
	lambda := func(a, b int) int {
		return a + b
	}
	fmt.Println(lambda(1, 2))
	fmt.Println(lambda(3, 2))
	fmt.Println(lambda(5, 2))
	fmt.Println(lambda(7, 2))
	fmt.Println(lambda(9, 2))

	func() {
		fmt.Println("Hello")
	}()

}
