package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "One"
			time.Sleep(500)
		}
	}()

	go func() {
		for {
			c2 <- "Two"
			time.Sleep(1000)
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println("REcieved", msg)

		case msg2 := <-c2:
			fmt.Println("REcieved", msg2)
		case <-time.After(1000):
			fmt.Println("timeout 2")

		}
	}

}
