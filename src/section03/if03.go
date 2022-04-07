//if문(3)

package main

import "fmt"

func main() {
	i := 100

	// if -else if 예제(1)
	if i >= 120 {
		fmt.Println("120이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("1")
	} else if i < 100 && i >= 50 {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
}
