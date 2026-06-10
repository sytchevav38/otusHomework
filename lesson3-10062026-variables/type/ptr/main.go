package main

import "fmt"

func main() {
	var a float64 = 3.4028235e+38

	fmt.Println(a)
	fmt.Println(&a)

	b := &a

	fmt.Println(b)
	*b = 2.3
	if b != nil {
		fmt.Println(*b)
	}
}
