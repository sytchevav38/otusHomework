package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi, Alex"

	fmt.Println(strings.Contains(str, "Alex"))
	fmt.Println(strings.Contains(str, "1"))
	fmt.Println(strings.HasPrefix(str, "Hi"))

	manyStrs := []string{
		"ads",
		"dsad",
		"dsd",
	}
	Joined := strings.Join(manyStrs, "-")
	fmt.Println(Joined)
	fmt.Println(strings.Split(Joined, "-"))
	strings.ReplaceAll(str, " ", "_")

}
