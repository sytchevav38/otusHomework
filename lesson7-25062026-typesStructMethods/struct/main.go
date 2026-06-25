package main

import "fmt"

// type user struct{}

// func main() {
// 	u := user{}

// 	fmt.Println(u)
// }

type user struct {
	Name     string
	Age      int
	lastName string
}

func main() {
	u := user{
		Name:     "John",
		Age:      20,
		lastName: "Ivanov",
	}

	fmt.Println("Старый чувак: ", u)
	u.Age = 33
	fmt.Println("Возраст: ", u.Age)
	u.Name = "Сеня"
	fmt.Println("Имя: ", u.Name)
}

func exampleLamdaStr() {
	var user struct {
		Name     string
		Age      int
		lastName string
	}

	user.Name = "Petya"
	user.Age = 13
	user.lastName = "Classov"

	fmt.Println(user)
}
