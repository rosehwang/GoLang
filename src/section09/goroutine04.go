//고루틴 기초04
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//고루틴
	//클로저 사용 예제
	//예제1

	runtime.GOMAXPROCS(1)

	s := "Goroutine Closure : "

	//for i := 0; i < 1000; i++ {
	for i := 0; i < 10; i++ {
		fmt.Println("test start")
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로저는 일반적으로 즉시 실행 But 고루틴 클로저는 가장 나중에 실행(반복문 종료 후 )
		fmt.Println("test end")
	}

	time.Sleep(5 * time.Second)
}
