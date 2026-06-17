package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	i := 0

	for i < 3 {
		i++
	}

	var users = []string{"a", "b", "c"}
	for i := 0; i < len(users); i++ {
		fmt.Println("users[i]", users[i])
	}

	var n int8 = 0
	for {
		println(n)
		n++
		if n == -1 {
			break
		}
	}
}
