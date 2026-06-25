package models

type (
	User struct {
		Name     string
		Age      int
		LastName string
	}
	Friend struct {
		Name     string
		Age      int
		LastName string
	}
)

func (u User) GetName() string {
	return u.Name
}
