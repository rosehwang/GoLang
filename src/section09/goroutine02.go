//고루틴 기초02
package main

import "fmt"
import "time"

func exe(name string) {
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>>", i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func main() {
	//고루틴
	exe("t1")
	//예제1
	fmt.Println("Main 고루틴 start :", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second) //
	fmt.Println("Main 고루틴 end :", time.Now())

	//
}
