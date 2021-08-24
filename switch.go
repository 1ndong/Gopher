package main

import "fmt"

func main() {
	a := 3
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a > 3")
	}

	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월 , 화")
	case "thrusday":
		fmt.Println("목")
	}

	temp := 18
	switch true {
	case temp < 10, temp > 30:
		fmt.Println("좋은날씨가 아닙니다")
	case temp > 10 && temp < 20:
		fmt.Println("추운날씨")
	case temp > 20 && temp < 30:
		fmt.Println("좋은날씨")
	}

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("teenage")
	case 33:
		fmt.Println("pair 3")
	default:
		fmt.Println("my age is", age)
	}

	fmt.Println("my favorite color is", colorToString(getMyFavoriteColor()))

	{
		a := 2
		switch a {
		case 1:
			fmt.Println("a == 1")
			break
		case 2:
			fmt.Println("a == 2")
			fallthrough
		case 3:
			fmt.Println("a == 3")
		}
		//when end of case moment, golang automatically call break internally
	}
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "red"
	case Blue:
		return "blue"
	case Green:
		return "green"
	case Yellow:
		return "yellow"
	default:
		return "undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func getMyAge() int {
	return 34
}
