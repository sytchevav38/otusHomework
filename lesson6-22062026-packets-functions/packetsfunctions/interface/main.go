package main

func main() {
	someFunc(1)
	someFunc("Строка")
	someFunc(1.6)
	someFunc('a')
}

func someFunc(a interface{}) {
}
