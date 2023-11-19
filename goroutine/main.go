package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		goroutine은 main function이 실행되는 동안에만 유효하다.
		메인 함수는 goroutine을 기다려주지 않는다.
	*/
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
