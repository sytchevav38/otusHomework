package main

import "fmt"

var firstString = "Hello world!"

func echoToTerminal(string string) {
	fmt.Println(string)
}

func main() {
	echoToTerminal(firstString)
}
