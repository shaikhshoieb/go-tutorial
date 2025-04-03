package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 1

}

func main() {
	fmt.Println("hello world")
	i := 100
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// The & operator gives the memory address of a variable.
	// The * operator gives the value at that address.
	// The * operator is also used to declare a pointer variable.
	// The & operator is used to get the address of a variable.
	// The * operator is used to dereference a pointer variable.

}
