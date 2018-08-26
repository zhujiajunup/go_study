package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Printf("numbers[1:4]==", numbers[1:4])
	fmt.Printf("numbers[:4]==", numbers[:4])
	fmt.Printf("numbers[1:]==", numbers[1:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len=%d\tcap=%d\tslice=%v\n", len(x), cap(x), x)
}
