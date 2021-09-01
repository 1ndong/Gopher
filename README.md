# Gopher

run via module  
&nbsp;&nbsp;&nbsp;&nbsp;go mod init [module name]  
&nbsp;&nbsp;&nbsp;&nbsp;go build  
&nbsp;&nbsp;&nbsp;&nbsp;./[module name].exe  

go pkg  
&nbsp;&nbsp;&nbsp;&nbsp;https://pkg.go.dev/  
&nbsp;&nbsp;&nbsp;&nbsp;https://pkg.go.dev/std -> std

https://vallista.kr/2019/12/28/%EB%8F%99%EC%8B%9C%EC%84%B1%EA%B3%BC-%EB%B3%91%EB%A0%AC%EC%84%B1-Concurrency-Parallelism
![image](https://user-images.githubusercontent.com/24688554/130925228-ae9fd23a-a5d8-40cc-aaa7-415eed7fb808.png)

1.workspace폴더에는 bin,pkg,src 3개의 서브폴더를 가져야 한다  
2.string타입의 zero value는 empty string이다  
3.import 검색 경로는 /src/로부터의 상대적 위치이다  
4.go에서 배열을 다른 변수에 할당하면, 배열 요소값 전체가 copy되어 별도의 배열이 된다  
5.배열의 크기는 type의 일부이다 [5]int와 [10]int는 다른 type이다  
6.c/c++과 다르게 1,0을 if문의 조건식으로 사용할 수 없다  
7.variadic function(가변인자함수)는 func f(s ...string)로 타입앞에 ...붙여줘야됨  
8.map[key]에 key를 발견하지 못하면 nil 혹은 zero value를 리턴한다  
9.interface{}는 empty interface로 모든 데이타 타입을 가질 수 있다  
{  
    var x interface{}  
    x = 1  
    x = "tom"  
    하면 x=1에서 x의 타입이 1로 결정될줄알앗는데  
    "tom" 할당하면 string으로 다시 값이 들어간다  
    void*같은 개념이므로 다시 값이 들어가는게 맞긴함  
    타입이 결정되는건 ts에서 let같은거였으면 말이 됨  
    
}  
10.
{
    package main
  
    import "fmt"
    
    func main() {
        c := make(chan int)
        c <- 1   
        fmt.Println(<-c) 
    }
    데드락이 걸린다. c채널에 1을 보내고 수신자를 기다리고 있는데 채널을 받는 수신자 고루틴이 없으므로 데드락이 발생  
}

11. 복수 채널을 기다리며 먼저 도착한 채널로 분기하기 위해 select문을 사용함
12. -> 좀더 봐야됨  
{
   func main() {
    ch := make(chan string, 1)
    f(ch)
    } 
    
    func f(ch chan<- string) {
        ch <- "OK"
        str := <-ch 
        println(str)
    } 
    수신 전용 채널 파라미터에 송신 <-을 사용하므로 컴파일 에러가 발생한다  
}

13.  
{
    func main() {
        var a interface{} = 1.0// -> 여기서 type이 float32로 지정됨
        i := a.(int)//type assertion이 발생함  
        fmt.Println(i)
    }
}

14.복수 라이브러리를 한 줄로 import하기 위해서는 import("fmt";"os")와 같이 표현할 수 있다  
15.go에서 concurrent 컴포넌트인 goroutine들은 채널으 ㄹ통해 메세지를 교환하면서 통신한다  
16.  
{
    func main() {
        for i := 0; i < 5; i++ {
            go func() {
                fmt.Println(i)
            }()
        }
    
        time.Sleep(1 * time.Second)
    }
    메인루틴이 실행이 다 되어서 i가 5가 되고 goroutine이 5번 실행되면서 i는 이미 5가 되었음 그래서 5가 5번 출력됨  
}  
17.go는inheritance대신 composition을 사용한다  
18.byte와 uint8 , rune와 int32는 동일한 타입을 가리킨다 하지만 int32는 int로 캐스팅될 수 없다 int는 시스템 비트에 따라서 int64가 될 수 있기 때문이다  
19.  
{  
    package main  
 
    func main() {  
        ch := make(chan int, 2)  
        ch <- 1  
        ch <- 2  
    
        close(ch)  
    
        println(<-ch)  
        println(<-ch)  
    }  
    close함수는 송신에 대해 채널을 닫지만, 수신은 허용한다 
    답은 1,2 가 차례로 출력된다   
}  
20.  
{
    func check(level int) string {
        if tag := "*"; level > 1000 {
            return string(level) + tag + tag + tag
        }
    
        return string(level) + tag
    }
    컴파일 에러가 발생한다 if문의 조건문 직전에 사용되는 optional statement는 if문 블럭 내에서만 사용된다  
}