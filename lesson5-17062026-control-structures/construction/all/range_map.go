package main

import "fmt"

func main() {
	//cats := []string{"Tom", "Jerry", "Flower"}
	catowners := map[string]string{
		"Tolya": "Tom",
		"Kolya": "Jerry",
		"Fedya": "Flower",
	}

	for owner, cat := range catowners {
		fmt.Println(owner, cat)
	}

	fmt.Println(catowners["Fedya"])

}
