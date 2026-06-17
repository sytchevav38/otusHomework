package main

import "fmt"

func main() {
LOOsP:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 3 && j == 7 {
				break LOOsP
			}
			// 	continue
			// }
			// if i == 7 {
			// 	break
			// }
			fmt.Println("i =", i, "j =", j)
		}
	}

}
