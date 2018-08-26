package main

func main() {
	var x, y int // 0, 0
	var (
		a int  // 0
		b bool // false
	)
	var c, d int = 1, 2 // 1, 2
	var e, f = 123, "hello"
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
