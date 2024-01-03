package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			<-ch1
			fmt.Print("a")
			ch2 <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			<-ch2
			fmt.Print("b")
			ch3 <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			<-ch3
			fmt.Print("c")
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	time.Sleep(20 * time.Second)
}
