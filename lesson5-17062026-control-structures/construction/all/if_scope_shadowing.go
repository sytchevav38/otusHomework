package shadowing

import "fmt"

func main() {
	x, y := 1, 2
	if true {
		x, y := 10, 20
		fmt.Println("внутри if (shadow):", x, y)
	}
	fmt.Println("после if (main):", x, y)
}
