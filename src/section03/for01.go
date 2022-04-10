//for01

package main

import "fmt"

func main() {
	//반복문 for
	//Go에서 유일하게 제공되는 반복문
	//다양한 사용법 숙지

	//예제1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	//에러발생1
	/*
		   for i := 0; i < 5; i++
				fmt.Println("ex1 : ", i)
	*/

	//예제2 (무한루프)
	/*
	 for {//무한루프 도는 반복문
	   fmt.Println("ex2 : hello Golang")
	 }
	*/

	//예제3 (Range 용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("Ex3: ", index, name)
	}
}
