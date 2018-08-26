package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Printf("Area of c1 = %f", c1.getArea())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func swap(x, y int) {
	var tmp int
	tmp = y
	y = x
	x = tmp
}

func swap2(x *int, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}
