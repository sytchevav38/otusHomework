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

func grade(score int) string {
	switch score {
	case 5:
		return "Отлично"
	case 4:
		return "Хорошо"
	case 3:
		return "Удовлитворительно"
	case 2, 1:
		return "пересдача"
	default:
		return "ошибка"
	}
}

func main() {
	for _, s := range []int{5, 4, 3, 2, 99} {
		fmt.Printf("score=%-3d > %s\n", s, grade(s))
	}
}
