package main

import "fmt"

type Book struct {
    Title  string
    Author string
}

// A function that updates the book title through a pointer
func ChangeTitle(b *Book, newTitle string) {
    b.Title = newTitle
}

func main() {
    // Create a new struct instance
    myBook := Book{
        Title:  "Go 101",
        Author: "John Doe",
    }

    // Create a pointer to 'myBook'
    bookPtr := &myBook

    fmt.Println("Before:", myBook)

    // Pass the pointer to the function
    ChangeTitle(bookPtr, "Go Mastery")

    fmt.Println("After: ", myBook)
}
