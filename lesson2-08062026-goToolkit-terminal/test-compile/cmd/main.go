package main

import (
	"fmt"
	pkg2 "my_module/internal/pkg2"
	pkg "my_module/pkg"
)

func main() {
	fmt.Println("Hello!")

	pkg.ASD()

	pkg2.ASD()
}
