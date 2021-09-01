package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("숫자를 입력하세요(종료:10)")
		var number int
		_, err := fmt.Scanln(&number) //_ intended not used
		if err != nil {
			fmt.Println("숫자가 아닙니다")

			stdin.ReadString('\n') //remove keyboard buffer(?)
			continue
		}
		fmt.Println("입력한 숫자는", number)
		if number == 10 {
			break
		}
	}

	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}

	{
		i := 1
		for {
			time.Sleep(time.Second)
			fmt.Println(i)
			i++

			if i > 3 {
				break
			}
		}
	}

	{
		a := 1
		b := 1

	OuterFor: //lable not recommended it can cause confusion(mess up instruct pointer)
		for ; a <= 9; a++ {
			for b = 1; b <= 9; b++ {
				if a*b == 45 {
					break OuterFor
				}
			}
		}
		fmt.Println(a, b)
	}
}
