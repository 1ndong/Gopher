package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	{
		s := make([]int, 5, 10) //type, len , capacity
		if s != nil {
			fmt.Println(len(s), cap(s))
		}
	}

	{
		var s []int
		if s == nil {
			fmt.Println(len(s), cap(s))
		}
	}

	{
		s := []int{0, 1, 2, 3, 4, 5}
		s = s[2:5] // result expect '[2,3,4]' first index inclusive , second index exclusive
		s = s[1:]  // 1 to end , [ : n] 0 to n-1
		fmt.Println(s)
	}

	{
		s := []int{0, 1}

		s = append(s, 2)       // 0, 1, 2
		s = append(s, 3, 4, 5) // 0,1,2,3,4,5

		fmt.Println(s)
	}

	{
		//like c++ std::vector,when over slice's capacity assign new slice(old slice's twice size) and copy whole data from old slice
		sliceA := make([]int, 0, 3)

		for i := 1; i <= 15; i++ {
			sliceA = append(sliceA, i)
			fmt.Println(len(sliceA), cap(sliceA))
		}

		fmt.Println(sliceA)
	}
	{
		sliceA := []int{1, 2, 3}
		sliceB := []int{4, 5, 6}

		sliceA = append(sliceA, sliceB...)
		//'...' means all element of sliceB
		//sliceA = append(sliceA, 4, 5, 6)

		fmt.Println(sliceA) // [1 2 3 4 5 6]
	}

	{
		source := []int{0, 1, 2}
		target := make([]int, len(source), cap(source)*2)
		copy(target, source)
		fmt.Println(target)               // [0 1 2 ]
		println(len(target), cap(target)) // 3, 6
	}
}
