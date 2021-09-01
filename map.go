package main

import "fmt"

func main() {

	var m map[int]string

	m = make(map[int]string)
	m[101] = "add"
	m[202] = "add2"

	str := m[101]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData)

	delete(m, 202)

	{
		tickers := map[string]string{
			"GOOD": "google inc",
			"MSFT": "microsoft",
			"FB":   "facebook",
		}

		val, exist := tickers["M4SFT"]
		if !exist {
			fmt.Println("not exist")
		} else {
			fmt.Println(val)
		}
	}

	{
		myMap := map[string]string{
			"A": "apple",
			"B": "banana",
			"C": "charlie",
		}

		//map is unordered
		for key, val := range myMap {
			fmt.Println(key, val)
		}
	}
}
