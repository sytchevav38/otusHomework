package iflesson

import "fmt"

func main() {
	x, y := 100, 200
	if x < y {
		fmt.Println("x меньше y")
	}
	if z := x * 2; z < y {
		fmt.Println("z<y, z =", z)
	}
	if err := mayFail(false); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Всё ок")
	}
}

func mayFail(fail bool) error {
	if fail {
		return fmt.error
	}
}
