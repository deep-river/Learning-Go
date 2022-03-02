/*
Go supports pointers, allowing you to pass references to values and records within your program
*/
package main

import "fmt"

func zeroval(ival int) {
	// Arguments passed to the function by value, zeroval will get a COPY of ival distinct from the one in the calling function
	ival = 0
}

func zeroptr(iptr *int) {
	// Function takes an int pointer, then DEREFERENCES the pointer from its memory address to the current value at that address. Assigning a value to a deferenced pointer changes the value at the referenced address.
	// *ADDRESS: dereference ADDRESS
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)
	// &VAR: get VAR's address
	fmt.Println("pointer: ", &i)
}