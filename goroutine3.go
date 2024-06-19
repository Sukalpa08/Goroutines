package main

import (
	"fmt"
	"sync"
	"time"
)

func go1() {
	fmt.Println("hhhhhh")

}

func go2(suku chan int) {
	fmt.Println("qqqqqqq", <-suku)

}

func go3(wg *sync.WaitGroup) {
	fmt.Println("zzzzzzz")
	wg.Done()

}

func main() {
	go go1()

	go func() {
		fmt.Println("wwwww")

	}()
	time.Sleep(1)

	var suku chan int = make(chan int)
	go go2(suku)
	suku <- 1

	var wg sync.WaitGroup
	wg.Add(1)
	go go3(&wg)
	wg.Wait()

}
