package main

import "fmt"

func describe(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("celoe chislo", v)
	case string:
		fmt.Println("Stroka:", v)
	case bool:
		fmt.Println("Logicheskiy tip", v)
	default:
	}
}

func main() {
	for _, s := range []int{5, 4, 3, 2, 99} {
		fmt.Printf("score=%-3d > %s\n", s, grade(s))
	}
}
