package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := "My name"

	// str += " Andrey"

	// fmt.Println(str)

	// for range 100 {
	// 	str += " the best!"
	// }

	// fmt.Println(str)

	str := "My name"

	stringBuilder := strings.Builder{}
	stringBuilder.Write([]byte(str))
	for range 100 {
		stringBuilder.Write([]byte(" Andrey!"))
	}

	str = stringBuilder.String()

	fmt.Println(str)

}
