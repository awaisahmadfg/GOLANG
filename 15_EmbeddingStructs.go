package main

import "fmt"

// Define a 'Person' struct
type Person struct {
    Name string
    Age  int
}

// Define an 'Employee' struct that embeds 'Person'
type Employee struct {
    Person
    ID     int
    Salary float64
}

func main() {
    // Initialize Employee with embedded fields
    emp := Employee{
        Person: Person{
            Name: "Dana",
            Age:  28,
        },
        ID:     201,
        Salary: 45000.0,
    }

    // We can directly access 'Name' or 'Age' because of embedding
    fmt.Println("Name:", emp.Name)
    fmt.Println("Age:", emp.Age)
    fmt.Println("ID:", emp.ID)
    fmt.Println("Salary:", emp.Salary)
}
