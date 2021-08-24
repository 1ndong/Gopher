package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다")
	} else if temp <= 18 {
		fmt.Println("나가자")
	} else {
		fmt.Println("덥다")
	}

	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}

	if filename, success := UploadFile(); success {
		if success {
			fmt.Println(filename)
		}
	}
}
func UploadFile() (filename string, success bool) {
	success = true
	filename = "filename"
	return
}

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}
