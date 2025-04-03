package main

import (
    "fmt"
)

type Employee struct {
    Name   string
    ID     int
    Salary float64
}

// Define a method to give a raise
func (e *Employee) GiveRaise(percentage float64) {
    // 'e' is a pointer to the Employee struct
    e.Salary += e.Salary * (percentage / 100)
}

func main() {
    emp := Employee{Name: "Charlie", ID: 103, Salary: 60000.0}
    fmt.Println("Before raise:", emp)

    // Call the GiveRaise method
    emp.GiveRaise(10) // Give a 10% raise
    fmt.Println("After  raise:", emp)
}

/*
Explanation
  1. We define a method GiveRaise(percentage float64) using a receiver func (e *Employee) ....
  2. The receiver (e *Employee) means GiveRaise can be called on an Employee pointer. Using a pointer receiver allows you to modify the original struct’s Salary.
  3. We update the Salary of the Employee instance emp.
  4. Finally, we demonstrate how to call the method and show the resulting updated Salary.
*/
