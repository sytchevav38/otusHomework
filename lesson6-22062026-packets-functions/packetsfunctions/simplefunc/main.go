package main

import "fmt"

func main() {
	fmt.Println(1, 23)

	sum(1, 2)

	//_ = sum(1,2)

	//sumvalue := sum(1, 2)

	//fmt.Println(sumvalue)

	_, multiplievalue := sumandmultiplie(1, 2)

	fmt.Println(multiplievalue)

}

func sum(a int, b int) int {
	return a + b
}

func sumandmultiplie(a, b int) (int, int) {
	return a + b, a * b
}
