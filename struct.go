package main

import "fmt"

type Person struct {
	name string
	age  int
}

type House struct {
	Address  string
	Size     int
	Price    int
	Category string
}

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User  // inheritance
	Price int
	Level int
}

type Test struct {
	a int32
	b float64
	//a 4byte b 8 -> unsafe.Sizeof(Test) -> 12x , 16(because data padding)
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
	{
		var house House = House{"서울시 서초구", 18, 50, "빌라"}
		fmt.Println(house)
	}
}
