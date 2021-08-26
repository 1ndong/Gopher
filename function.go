package main

import "fmt"

func main() {
	fmt.Println(Add(3, 5))
	fmt.Println(Divide2(1, 0))
	fmt.Println(Divide2(4, 2))

	{
		msg := "hello"
		say(&msg)
		fmt.Println(msg)

		say2("33", "44", "6^", "&^", "**")

		sum := func(n ...int) int {
			s := 0
			for _, t := range n {
				s += t
			}
			return s
		}

		result := sum(1, 2, 3, 4, 5)
		fmt.Println(result)
	}

	{
		add := func(i, j int) (result int) {
			result = i + j
			return
		}

		r1 := calc(add, 10, 20)
		fmt.Println(r1)

		r2 := calc(func(x, y int) int { return x - y }, 10, 20)
		fmt.Println(r2)

		r3 := calc2(add, 30, 50)
		fmt.Println(r3)
	}
}

type calculator func(int, int) int

func calc2(f calculator, a, b int) int {
	result := f(a, b)
	return result
}

func calc(f func(int, int) int, a, b int) int {
	result := f(a, b)
	return result
}

func say(msg *string) {
	fmt.Println(*msg)
	fmt.Println(msg)
	*msg = "changed"
}

func say2(msg ...string) {
	for _, s := range msg {
		fmt.Println(s)
	}
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
