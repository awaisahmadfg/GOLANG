package main

import "fmt"

func sumNumbers(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    result := sumNumbers(1, 2, 3, 4, 5)
    fmt.Println(result) // 15
    
    // You can also pass a slice to a variadic function by using the ... operator:
    nums := []int{10, 20, 30}
    result = sumNumbers(nums...) // "Unpacking" the slice into the function
    fmt.Println(result)           // 60
}

/* 
  Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
*/
