////상수1 - 자바로 치면 static final
package main

import "fmt"

func main() {
	const a, b int = 10, 90
	const c, d, e = true, 0.84, "test"
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 7776
	)
	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)
}
