package initpkg

import "fmt"

var (
	a = b + c
	b = f()
	c = f()
	d = 3
)

func init() {
	fmt.Println("init")
}

func f() int {
	d++
	fmt.Println(d)
	return d
}
