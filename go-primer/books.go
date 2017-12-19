package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go语言"
	Book1.author = "dwong"
	Book1.subject = "Go语言基础"
	Book1.book_id = 6495700

	Book2.title = "python语言"
	Book2.author = "dwong"
	Book2.subject = "python语言基础"
	Book2.book_id = 6495704

	//	fmt.Printf("Book 1 title : %s\n", Book1.title)
	//	fmt.Printf("Book 1 author : %s\n", Book1.author)
	//	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	//	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)
	//
	//	fmt.Printf("Book 2 title : %s\n", Book2.title)
	//	fmt.Printf("Book 2 author : %s\n", Book2.author)
	//	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	//  fmt.Printf("Book book_id : %d\n", Book.book_id)
	printBook(Book1)
	printBook(Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
