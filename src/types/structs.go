package main

// Structs are collection of fields
// They’re useful for grouping data together to form records
import "fmt"

type Vertex struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// struct fields are accessed using a dot
	v := Vertex{1, 2}
	v.x = 4
	fmt.Println(v.x)

	// example 2
	fmt.Println(person{"Fred", 23})
	s := person{name: "Bob", age: 39}
	fmt.Println(s.name)
	// Struct fields can be accessed through a struct pointer.
	sp := &s
	fmt.Println(sp.name)
	fmt.Println(sp)
}
