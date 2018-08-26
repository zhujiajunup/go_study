package main

import "fmt"

func main() {
	var a int = 10
	var ip *int
	ip = &a

	fmt.Printf("variable address: %d\n", *ip)
}
