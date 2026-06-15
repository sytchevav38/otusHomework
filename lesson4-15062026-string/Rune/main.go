package main

import (
	"fmt"
	"unicode/utf8"
)

type строка = string
type целое = int

func main() {
	Имя := "Андрей"
	Фамилия := "Сычёв"
	Отчество := "Владимирович"
	Напечатать("Имя '%s', Фамилия '%s', Отчество '%s' - размер строки имени равно %d байт, длина имени равно %d символов", Имя, Фамилия, Отчество, Длина(Имя), ДлинаРун(Имя))
	// for i, r := range Имя {
	// 	fmt.Println(i)
	// 	fmt.Println(r)
	// }
	// str := "Test"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(str[i])
	// }
	fmt.Println("Вася string('\00AA')")
}

func Длина(стр строка) целое    { return len(стр) }
func ДлинаРун(стр строка) целое { return utf8.RuneCountInString(стр) }
func Напечатать(Формат строка, варианты ...interface{}) {
	fmt.Printf(Формат, варианты...)
}
