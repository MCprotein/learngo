package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		goroutine은 main function이 실행되는 동안에만 유효하다.
		메인 함수는 goroutine을 기다려주지 않는다.

		channel: gouroutine과 메인함수, 혹은 다른 goroutine 사이에 정보를 전달하기 위한 방법

	*/
	channel := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, channel)
	}

	result := <-channel
	fmt.Println(result)

	fmt.Println(<-channel)

}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	channel <- true
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
