/*
	Go's struct are typed collection of fields. They're useful for the grouping data together to form records.
*/

package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

// NewPerson constructs a person struct with the given name
func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 23

	// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	return &p
}

func main() {
	// This syntax create a new struct
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initialing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions.
	fmt.Println(NewPerson("John"))

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can able use dots with struct pointers. The pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Struct are mutable.
	sp.age = 51
	fmt.Println(sp.age)
}
