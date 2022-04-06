//if문(1)

package main

import "fmt"

func main() {
	//제어문(조건문)
	// IF문 : 반드시 Boolean 형 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
	//소괄호 사용 X

	var a int = 20 //선언된 변수가 프로그램내 사용되지 않는다면, 에러를 발생시킴. 따라서 사용 X변수는 삭제해야함.
	//변수를 선언하면서 초기값을 지정하지 않으면, Zero Value를 할당한다. 숫자형에는 0, bool은 false, string ""
	b := 20 //변수를 선언하는 또 다른방식. Short Assignment Statement (:=) 사용가능. 함수(func)내에서만 사용가능
	//함수 밖에서는 var 을 사용해야한다.
	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	//에러 발생 1
	/*
	  if b>= 25
	  {//여기다 괄호를 넣어버리면 실행될때 b>=15에 세미콜론(문장끝)으로 인식해버려서 에러남
	    fmt.Println("25이상")
	  }
	*/
	//에러 발생2
	/* //괄호 또 안쓰면 에러남
	   if b>=25
	     fmt.Println("25이상")
	*/
	//에러발생3
	/*
	  if c:=1; c {

	  }
	*/

	//에러발생4
	/*
		  if c := 40; c >= 35 {
				fmt.println("35이상")
			}

			c += 20
	*/
}
