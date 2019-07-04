package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("timer started")
	now := time.Now()

	duration := 3 * time.Second
	<-time.After(duration)

	fmt.Println("timer ended")
	fmt.Printf("elapsed time: %.3f\n", time.Now().Sub(now).Seconds())
}
