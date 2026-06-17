package main

import "fmt"

func main() {
	size := 25
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if (row+col)%2 == 0 && col == size-1 {
				fmt.Print(" \n")
			} else if (row+col)%2 == 0 {
				fmt.Print(" ")
			} else if col == size-1 {
				fmt.Print("#\n")
			} else {
				fmt.Print("#")
			}
		}

	}
}
