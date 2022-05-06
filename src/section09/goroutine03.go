//고루틴 기초03
package main

import "fmt"
import "time"
import "math/rand"
import "runtime"

func exe(name int) {
	r := rand.Intn(100) //100미만의 랜덤한 숫자를 가져온다.
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func main() {
	//고루틴
	//멀티 코어(다중 cpu) 최대한 활용
	runtime.GOMAXPROCS(runtime.NumCPU()) //내 현재 CPU의 물리코어 갯수
	// 이 코어를 다 써서 실행한다.

	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0))

	//예제1
	fmt.Println(" Main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i) //고루틴 100개 생성
	}
	fmt.Println(" Main Routine End : ", time.Now())
}
