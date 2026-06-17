package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 || i == 4 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println("i =", i)
	}

	users := []string{"a", "b", "c"}
	for i := 0; i < len(users); i++ {
		if users[i].hasSubscriprion() {
			continue
		}
		users[i].sendmessage()
	}
	for _, user := range users {
		if user.hasSubscriprion() {
			continue
		}
		user.sendmessage()
	}

}
