package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		i := i
		defer func() {
			fmt.Println(i)
		}()
	}
}
