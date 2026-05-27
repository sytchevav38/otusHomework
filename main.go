package main

import "fmt"

var firstString = "Hello world!"

/*
func echoToTerminal(string string) {
	fmt.Println(string)
}
*/

func main() {
	var name string
	fmt.Println("Привет, как тебя зовут?")
	fmt.Scanln(&name)
	fmt.Println("Я понял, тебя зовут:", name)
}
