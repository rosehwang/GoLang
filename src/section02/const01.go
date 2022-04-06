//상수1 - 자바로 치면 static final
package main

import "fmt"

func main() {
	//상수
	//const 사용 초기화, 한 번 선언후 값 변경 금지, 고정된 값 관리용
	//선언과 동시에 초기화 해아함
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	//const d = getHeight() 함수 사용 불가능
	const e = 35.6
	const f = false
	/*
	  에러발생
	  const a string
	  g = "Test3"
	*/
	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)

}
