package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	if d <= 0 {
		return
	}

	<-time.After(d)
}

func main() {
	fmt.Println("start:", time.Now())
	sleep(1500 * time.Millisecond)
	fmt.Println("done :", time.Now())
}
