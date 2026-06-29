package main

import "fmt"

func main() {
	mm := map[string]string{
		"yellow": "жёлтый",
		"green":  "зелёный",
		"black":  "чёрный",
		"red":    "красный",
	}

	for k, v := range mm {
		fmt.Printf(
			"Ключ: %s | Значение: %s \n",
			k,
			v,
		)
	}

	fmt.Println(mm)

	_, ok := mm["rose"]
	fmt.Println(ok)

	delete(mm, "green")

	for k, v := range mm {
		fmt.Printf(
			"Ключ: %s | Значение: %s \n",
			k,
			v,
		)
	}

	_, ok = mm["rose"]
	fmt.Println(ok)

	mm["red"] = "КРАСИВЫЙ"

	for k, v := range mm {
		fmt.Printf(
			"Ключ: %s | Значение: %s \n",
			k,
			v,
		)
	}

}
