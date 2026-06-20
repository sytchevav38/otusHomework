// package main

// import "fmt"

// func main() {
// 	size := 25 //размер доски - по сути сторона квадрата, т.е. значение и на ширину и на высоту влияет.
// 	for row := 0; row < size; row++ {
// 		for col := 0; col < size; col++ {
// 			if (row+col)%2 == 0 && col == size-1 {
// 				fmt.Print(" \n")
// 			} else if (row+col)%2 == 0 {
// 				fmt.Print(" ")
// 			} else if col == size-1 {
// 				fmt.Print("#\n")
// 			} else {
// 				fmt.Print("#")
// 			}
// 		}

//		}
//	}

package main

import (
	"flag"
	"fmt"
)

func main() {
	size := flag.Int("size", 8, "Размер доски") //Принимает значение из консоли, go run main.go -size int 15, если значение не передаётся, size принимает значение 8 по умолчанию.
	flag.Parse()
	for row := 0; row < *size; row++ {
		for col := 0; col < *size; col++ {
			if (row+col)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
