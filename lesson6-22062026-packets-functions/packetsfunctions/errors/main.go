package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err error
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Пользователь делит на ноль: ", err.Error())
	} else {
		fmt.Println(result)
	}
	// fmt.Println(divide(1, 0))
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Ну и зачем ты делишь на 0? В школе не учился чтоли?!")
	}
	return a / b, nil
}
