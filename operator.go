package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

	c := 10
	c++ //this is ok
	fmt.Println(c)
	//fmt.Println(c++)
	//-> fail , increment is ok but increment and assignment fail
	//like b = c++

	/*
		operator priority(5 is first)
		5 * / % << >> & &^
		4 + - | ^
		3 == != < <= > >=
		2 &&
		1 ||
	*/
}
