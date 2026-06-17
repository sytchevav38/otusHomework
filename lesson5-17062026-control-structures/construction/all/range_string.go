package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Привет"
	fmt.Println("len(s):", len(s))
	fmt.Println("Rune:", utf8.RuneCountInString(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("i=%d byte=%d\n", i, s[i])
	}

	for i, r := range s {
		fmt.Printf("byteIndex=%d rune=%c (%d)\n", i, r, r)
	}

}
