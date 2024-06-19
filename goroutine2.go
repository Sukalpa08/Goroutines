package main

import (
	"fmt"
)

func routine(ch chan int) {

	for i := range ch {

		if i == 1 {
			fmt.Println("Direct validation")
		}
		fmt.Println("channel library function", i)

	}
}

func main() {

	var channel chan int = make(chan int)
	go routine(channel)

	channel <- 1
}
