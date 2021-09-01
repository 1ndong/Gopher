package testlib

import "fmt"

//import _ 'package name'-> '_' means only call init()function concurrently import moment

var pop map[string]string

func init() {
	pop = make(map[string]string)
	fmt.Println("init")
}

func Lib() {
	fmt.Println("lib")
}
