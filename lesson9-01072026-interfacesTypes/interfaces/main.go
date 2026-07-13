package main

import (
	"fmt"
	repo "interfaces/repo"
)

type UserRepo interface {
	CreateUser(id int, name string) error
	UpdateUser(id int, name string) error
	DeleteUser(id int) error
}

func handleCreateUser() {
	ur := repo.NewUserRepo()

	err := ur.CreateUser(1, "Alex")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var f float64 = 123.456
	i := int(f)  // Преобразование одного типа данных к другому
	b := f.(int) // Приведение одного типа данных к другому, только попытка, не факт что сработает.

	fmt.Println(i) // 123
	fmt.Println(b)
	fmt.Printf()
}
