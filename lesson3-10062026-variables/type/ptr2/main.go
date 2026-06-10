package main

import "fmt"

func main() {
	a := new(int) //Указатель на чё-то
	fmt.Println(*a)
	*a = 10
	fmt.Println(*a)
}
