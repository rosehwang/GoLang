package main

import "fmt"

func main() {
	fmt.Println("hello world")
} //ctrl + shift + B

/**
go run .\helloworld.go
메모리에서 빌드를 하지않고 바로 실행하는 것(테스트개념)
*/

/**
go build .\helloworld.go
.\helloworld.go
해당 폴더에 .exe 실행 파일 생성
실행가능한 파일을 만드는 것(테스트 개념)
*/

/**
go install
개발을 다 완료했을 때 해당 연관 파일 및 프로젝트 로 빌드해서
폴더명.exe 파일을 생성.
*/
