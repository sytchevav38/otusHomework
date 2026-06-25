package models

import "fmt"

type User struct {
	Name     string
	Age      int
	LastName string
}

// type Employee struct {
// 	CompanyName string
// 	YearCount   int
// 	UserInfo    User
// }

// func a() {
// 	e := Employee{
// 		CompanyName: "Yandex",
// 		YearCount:   10,
// 		UserInfo: User{
// 			Name:     "Test",
// 			Age:      999,
// 			LastName: "Testov",
// 		},
// 	}

// 	fmt.Println(e.UserInfo.Name)
// }

type Employee struct {
	CompanyName string
	YearCount   int
	User
}

func a() {
	e := Employee{
		CompanyName: "Yandex",
		YearCount:   10,
		User: User{
			Name:     "Test",
			Age:      999,
			LastName: "Testov",
		},
	}

	fmt.Println(e.User.Name)
}
