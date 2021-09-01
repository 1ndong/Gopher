package gettest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTest(){
	resp , err := http.Get("https://httpbin.org/get")
	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()

	data, err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(data))
}