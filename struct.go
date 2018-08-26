package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var go_in_action Book
	go_in_action.title = "Go in Action"
	go_in_action.author = "William Kennedy"
	go_in_action.subject = "Go in Action introduces the Go language, guiding you from inquisitive developer to Go guru."
	go_in_action.book_id = 1
	fmt.Println(go_in_action)
}
