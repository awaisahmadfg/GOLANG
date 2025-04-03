package main

import (
    "fmt"
)

// Define the Employee struct
type Employee struct {
    Name   string
    ID     int
    Salary float64
}

func main() {
    // Create an instance of Employee with initial values
    emp1 := Employee{
        Name:   "Alice",
        ID:     101,
        Salary: 50000.0,
    }

    // Another way: create an empty instance, then assign
    var emp2 Employee
    emp2.Name = "Bob"
    emp2.ID = 102
    emp2.Salary = 55000.0

    // Print struct values
    fmt.Println("Employee 1:", emp1)
    fmt.Println("Employee 2:", emp2)
}
