package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "hello world"

	a = 20
	msg = "good morning"
	fmt.Println(msg, a)

	type myInt int
	var aa myInt = 20
	fmt.Println(aa)

	bb := 30
	fmt.Println(bb)

	var cc int
	fmt.Println(cc)

	{
		a := 3
		var b float64 = 3.5
		var c int = int(b)
		d := float64(a) * b
		var e int64 = 7
		f := a * int(e)
		//a is int64 at 64bit pc e is int64 but disallowed
		//because a is able to inference to int32 at 32bit pc

		fmt.Println(a, b, c, d, e, f)
	}

	fmt.Printf("%d\n", 1)
	/*
		%T -> print data type
		%t -> bool
		%v -> adjust data type
		%b -> binary
		%c -> unicode character
		%o -> octa
		%O -> octa add 'Oo' which means octa value
		%x -> hex over 10 -> 'a-f'
		%X -> hex over 10 -> 'A-F'
		%q -> print string ignore special character
	*/
	{
		var a string
		n, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n, a)
		}
	}
}

/*
byte == uint8 1byte unsigned int
rune -> 1char, n*rune -> string , utf8 1~3byte -> int32 4byte
int -> 32bit pc int32 , 64bit pc int64

bool
string
array -> [size]
slice -> [] not assigned size

type -> typedef at c++

var a int = 10
var a int -> init value is 0
var a = 10 -> type inference
a:=10

default float is float64

init value
int -> 0
float -> 0.0
bool -> false
string -> "" empty string
else -> nil

different type operation disallowed
want to avoid

*/
