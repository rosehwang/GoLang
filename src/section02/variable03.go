//변수1
package main

import "fmt"

func main() {
	//짧은 선언
	//함수 안에서만 사용(전역x), 선언 후 할당 예외 발생
	//주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있다 -> 추천
	//자바로 치면 지역변수 느낌인가 ㅋ
	shortVal1 := 3
	shortVal2 := "Test"
	shortVal3 := false
	fmt.Println("shortVal1 :", shortVal1, "shortVal2 :", shortVal2, "shortVal3 :", shortVal3)
}
