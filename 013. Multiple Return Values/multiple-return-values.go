/*
	Go has built-in support for multiple return values. This feature is used often in idiomatic GO, for example to return both result and error value from a function

	Note:
	Accepting a variable number of arguments is another nice feature of Go functions.
*/

package main

import "fmt"

// The (int, int) in this function signature shows that the function return 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Here we use the 2 different return values from the call with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if you only want a subset of the returned values, use the black identifier
	_, c := vals()
	fmt.Println(c)
}
