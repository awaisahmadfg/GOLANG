package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print message to the user
	fmt.Println("Enter a sentence:")

	// Create a reader to read input from the user
	reader := bufio.NewReader(os.Stdin)

	// Read the full line of text entered by the user
	text, _ := reader.ReadString('\n')

	// Split the sentence into words using spaces
	words := strings.Fields(text)

	// Count the number of words
	count := len(words)

	// Print the result
	fmt.Println("Word count:", count)
}
