package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(1)
	//go fmt.Println(100)
	for i := range 100 {
		//fmt.Println(i)
		go fmt.Println(i) // горутина
	}

	time.Sleep(time.Second)
}
