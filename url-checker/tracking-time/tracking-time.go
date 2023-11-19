package trackingtime

import (
	"fmt"
	"time"
)

func Trackingtime() func() {
	startTime := time.Now()
	return func() {
		elapsedTime := time.Since(startTime)
		fmt.Println("실행 시간", elapsedTime)
	}
}
