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


/*USE CASES:
A practical reason to use pointers in Go is whenever you need to modify or work with the actual value of something rather than a copy. Here are a few everyday use cases:

1.  Modifying a value in a function:
    If you want a function to directly change a variable’s value from outside that function (like updating a field in a struct), passing a pointer is essential. Otherwise, Go passes your variable by value (as a copy), so changes won’t affect the original.

2.  Improving performance for large data:
    When you have a large struct or array, passing it around by value can be expensive because Go copies it each time. By passing a pointer (*MyStruct), you pass only the address, which is much lighter. This can prevent excessive memory use or slowdowns.

3.  Sharing a single instance:
    Sometimes you want multiple parts of your program to refer to and update the same piece of data. Pointers are how you achieve that shared, “single source of truth.”

4.  Linked data structures:
    Data structures like linked lists, trees, and other references-based structures require pointers. Each node holds a pointer to its children or next/previous nodes.


*/
