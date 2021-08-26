package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{}

	p.name = "kim"
	p.age = 33

	fmt.Println(p)

	{
		p := new(Person) // pointer
		p.name = "nam"
	}
}
