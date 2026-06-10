package main

import "fmt"

func main() {
	var (
		ba, ab = 1, 2
	)
	const b = 22
	var a int = 10
	c := 23
	//var d float64 = 3.25
	a = 15
	var bol bool = true
	bol = false
	var max int8 = 127
	//var test float32

	fmt.Println(a + c)
	fmt.Println(ab + ba)
	fmt.Println(bol)
	fmt.Println(max)
	max += 1
	fmt.Println(max)
}
