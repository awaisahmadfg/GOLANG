/* Passing by value: (zeroval) uses a copy, so it can’t alter the original variable in Main.
 Passing by pointer: (zeroptr) gets direct access to the variable’s memory address, so it can change the original value inside main.
*/

package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i) // Passing copy, the value will remain same that was in main 
    fmt.Println("zeroval:", i)

    zeroptr(&i) // This will change the value becasue That will sets the actual i in memory to 0 where that variable memory address exist.
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}
