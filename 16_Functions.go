package main

import (
    "fmt"
)

func greet(name string) string {
    return "Hello, " + name + "!"
}

func main() {
    message := greet("Satoshi")
    fmt.Println(message)
}
