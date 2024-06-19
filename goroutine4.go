package main

import "fmt"

func sync(b chan int, b1 chan int) {

	select {

	case c := <-b:
		fmt.Println("Popcorn", c)

	case d := <-b1:
		fmt.Println("Coke", d)
	}
	fmt.Println("Property division over")

}

func main() {
	var buf = make(chan int)
	var buf1 = make(chan int)
	go sync(buf, buf1)

	buf1 <- 2
	fmt.Println("gooooooogleeeee")

}
