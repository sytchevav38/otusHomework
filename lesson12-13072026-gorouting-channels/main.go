package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(1)
	//go fmt.Println(100)
	//ch := make(chan int, 10)
	for i := range 100 {
		//fmt.Println(i)
		go func() {
			fmt.Println(i) //в таком формате может быть выведено старое значение дважды
		}()
		//go fmt.Println(i) // горутина
		go func(test string) { fmt.Println(test) }("Тест")
	}

	time.Sleep(time.Second)
}
