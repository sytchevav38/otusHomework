package main

const PI = 3            // принимает подходящий тип
const pi float32 = 3.14 // пстрогий тип
const (
	TheA = 1
	TheB = 2
)

const (
	X = iota // 0
	Y        // 1
	Z        // 2
)

const (
	k = 3
)

var (
	d int
	e = 1
)

var m = 1
var n int

func main() {
	var a int

	b := 1

	const c = 3
	println(d + e + k)
	println(a + b + c)
}
