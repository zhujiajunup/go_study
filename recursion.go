package main

import "fmt"

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

func main() {
	i := 4
	fmt.Printf("%d! = %d", i, Factorial(i))
}
