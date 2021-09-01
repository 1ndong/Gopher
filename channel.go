package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	fmt.Println(i)

	{
		done := make(chan bool)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
			}
			done <- true
		}()

		fmt.Println(<-done)
	}

	{
		c := make(chan int, 1)
		c <- 1
		fmt.Println(<-c)
	}

	//'channel name' <- is send only chan
	//<- 'channel name' is receive only chan
	{
		ch := make(chan string, 1)
		sendChan(ch)
		receiveChan(ch)
	}

	{
		ch := make(chan int, 2)
		ch <- 111
		ch <- 222
		close(ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		if _, success := <-ch; !success {
			fmt.Println("not more data")
		}
	}

	{
		ch := make(chan int, 2)

		ch <- 123
		ch <- 456
		close(ch)
		/* loop until channel is empty
		for {
			if i, success := <-ch; success {
				fmt.Println(i)
			} else {
				break
			}
		}
		*/
		for i := range ch {
			fmt.Println(i)
		}
	}

	{
		done1 := make(chan bool)
		done2 := make(chan bool)
		go run1(done1)
		go run2(done2)

	EXIT:
		for {
			select {
			case <-done1:
				fmt.Println("run1 complete")
			case <-done2:
				fmt.Println("run2 complete")
				break EXIT
			}
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func sendChan(ch chan<- string) {
	ch <- "data"
	//x:= <-ch is error because ch is send only channel
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
