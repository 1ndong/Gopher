package main

import (
	"fmt"
)

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	anotherNext := nextValue()
	fmt.Println(anotherNext()) // 1 다시 시작
	fmt.Println(anotherNext()) // 2
}
