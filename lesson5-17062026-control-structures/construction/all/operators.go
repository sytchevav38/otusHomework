package operators

import "fmt"

func main() {
	x, y := 5, 9
	fmt.Println("x < y :", x < y)
	fmt.Println("x == y :", x == y)
	fmt.Println("x != y :", x != y)
	fmt.Println("x > y :", x > y)

	a, b := true, false
	fmt.Println("a && b :", a && b)
	fmt.Println("a || b :", a || b)
	fmt.Println("!a :", !a)

	var p *int
	if p != nil && *p > 0 {
		fmt.Println("Положительное")
	} else {
		fmt.Println("p == nil, правая часть не вычислялась")
	}
	var user *string
	user, err := getUser(id)
	if err != nil {
		//....
	}

	if user != nil && user.hasSubscriprion() {
		// show banner
	}
}
