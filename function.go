package main

import "fmt"

func main() {
	fmt.Println(Add(3, 5))
	fmt.Println(Divide2(1, 0))
	fmt.Println(Divide2(4, 2))
}

//Add -> if first character is uppercase this function is public access modifier
//lower case means private access modifier
func Add(a int, b int) int {
	return a + b
}

//multi return!
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}
