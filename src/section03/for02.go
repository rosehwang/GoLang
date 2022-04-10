//for02

package main

import "fmt"

func main() {
	//예제1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}

	fmt.Println("Ex1 : ", sum1)

	//예제2 - 가독성이 좋음 ㅋ
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		//j := i++  //GO에서 후치연산은 반환 값 X
	}
	fmt.Println("Ex2 : ", sum2)

	//예제3 - while문이랑 비슷하게 사용하는 것
	sum3, i := 0, 0

	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("Ex3 : ", sum3)

	//예제4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 { // i++, j += 10 이런식으로 쓰면 안됨
		fmt.Println("Ex4 : ", i, j)
	}

}
