package main

import "fmt"

func scoringResult(score int) string {
	if score > 90 {
		return "good"
	} else if score > 75 {
		return "not bad"
	} else if score > 60 {
		return "bad"
	}
	return "fool"
}

func main() {
	for _, s := range []int{95, 80, 65, 40} {
		fmt.Printf("score=%-3d > %s\n", s, scoringResult(s))
	}
}
