//고루틴 기초05(고루틴 내용 추가)
package main

import (
  "fmt"
  "sync"
  "runtime"
)

func main() {
//02. 익명 함수 Goroutine
//고루틴은 익명함수에 대해 사용할 수 있다. 즉 go 키워드 뒤에
//익명함수를 바로 정의하는 것으로, 이는 익명함수를 비동기로 실행하게 된다.
//고루틴으로 실행하면 비동기적으로 그 익명함수를 실행하게 됨.
//두번째 익명함수는 파라미터를 전달하는 예제로 익명함쉐 파라미터가 있는 경우,
//go 익명함수 바로 뒤에 파라미터를 하계 전달하기 된다.

   // WaitGroup 생성. 2개의 Go루틴을 기다림.
   /*
   var wait sync.WaitGroup
   wait.Add(2)

   // 익명함수를 사용한 goroutine
   go func() {
       defer wait.Done() //끝나면 .Done() 호출
       fmt.Println("Hello")
   }()

   // 익명함수에 파라미터 전달
   go func(msg string) {
       defer wait.Done() //끝나면 .Done() 호출
       fmt.Println(msg)
   }("Hi")

   wait.Wait() //Go루틴 모두 끝날 때까지 대기
   */

   //03. 다중 CPU 처리
   // Go는 디폴트로 1개의 CPU사용. 즉 여러개의 고루틴을 만들더라도 1개의 CPU에서
   //작업을 시분활하여 처리한다.(Concurrent 처리). 만약 머신이 복수개의 CPU를 가진경우,
   //Go프로그램을 다중 CPU에서 병령처리하게 할 수 있는데, 병렬처리를 위해서는 아래와 같이
   //GOMAXPROCS(CPU수)함수를 호출하여야 한다.

   // 4개의 CPU 사용
   runtime.GOMAXPROCS(4)

   var wait sync.WaitGroup
   wait.Add(2)

   // 익명함수를 사용한 goroutine
   go func() {
       defer wait.Done() //끝나면 .Done() 호출
       fmt.Println("Hello")
   }()

   // 익명함수에 파라미터 전달
   go func(msg string) {
       defer wait.Done() //끝나면 .Done() 호출
       fmt.Println(msg)
   }("Hi")

   wait.Wait() //Go루틴 모두 끝날 때까지 대기
}
