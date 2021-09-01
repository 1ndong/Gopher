package main

import "fmt"

func main() {
	var ary [10]int
	ary[0] = 10
	ary[2] = 30
	fmt.Println(ary)

	var t [5]float64 = [5]float64{1.3, 2.0, 3.0, 4.0, 5.0}
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}

	days := [3]string{"monday", "tuesday", "wednesday"}
	fmt.Print(days)
	var temps [5]float64 = [5]float64{1.4, 2.5}
	fmt.Print(temps)
	var s = [5]int{1: 10, 3: 30}
	fmt.Print(s)
	x := [...]int{10, 20, 30} //[3]int
	//[]int{10,20,30} -> slice(dynamic size)
	fmt.Print(x)
}
