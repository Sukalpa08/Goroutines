package main

import (
	"fmt"
	"time"
)

func Funtion() {

	fmt.Println("Stack memory")
}

func goroutine() {

	fmt.Println("Special function")
}

func main() {
	Funtion()
	go goroutine() //Inorder to execute goroutine function we execute library functions

	time.Sleep(1)

	func() {

		fmt.Println("OS memory")

	}()
}
