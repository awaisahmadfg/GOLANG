/*
  Methods vs Function:
      1. A method is essentially a function, but it is associated with a specific type (usually a struct). 
      2. Methods are functions that have a "receiver" argument. This receiver allows the method to access and manipulate the data of the type it is associated with.
      3. Receiver: The receiver is like an argument to the method, but it binds the method to a specific type. It can be a value receiver or a pointer receiver.
*/

package main

import "fmt"

type Rectangle struct {
    width, height int
}

// Method with value receiver (does not modify the original struct)
func (r Rectangle) area() int {
    return r.width * r.height
}

// Method with pointer receiver (modifies the original struct)
func (r *Rectangle) Rarea() int {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    result := rect.area() // Calling the method on the instance
    fmt.Println(result)   // Outputs: 50
    
    result = rect.Rarea()
    fmt.Println(result)
}
