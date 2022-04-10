//for03

package main

import "fmt"

func main() {
	//예제1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1 //해당 조건문을 충족할 때 break 하게 됨.
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

	//예제2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //바로 다음 i로 넘어감 아래 예문 진행이 안됨
		}
		fmt.Println("ex2 : ", i)
	}

	//예제3
Loop2:
	//에러 발생 (Loop 레이블 밑에 관련이 없는 소스 코드 )
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2 //continue면 아래 예문을 실행하지 않고 바로 다음 i번째가 진행이 된다.
			}

			fmt.Println("ex3 : ", i, j)
		}
	}
}
