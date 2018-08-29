package main

import "fmt"

func main() {
	threeIndexSlice()
}

func threeIndexSlice() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	threeIdxS := source[2:3:4]
	printStringSlice(threeIdxS)
}

func sliceTest() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	printSlice(newSlice)
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

func printStringSlice(s []string) {
	fmt.Printf("len=%d\tcap=%d\tslice=%v\n", len(s), cap(s), s)
}
