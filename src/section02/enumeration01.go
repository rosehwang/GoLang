//열거형1
package main

import "fmt"

func main() {
	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음

	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C) // 0 1 2
}
