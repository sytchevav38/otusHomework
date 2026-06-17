package variables

import "fmt"

func main() {
	var a int
	fmt.Println("var a int	->", a)
	a = 9
	fmt.Println("a = 9	->", a)
	b := 7
	fmt.Printf("b:=7	-> %d (тип %T)\n", b, b)

	x, err := half(10)
	fmt.Println("x, err:=	->", x, err)
	y, err := half(20)
	fmt.Println("y, err:=	->", y, err)
}

func half(n int) (int, error) {
	return n / 2, nil
}
