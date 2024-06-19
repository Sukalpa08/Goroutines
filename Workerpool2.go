package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, result chan<- int) {

	for task := range tasks {
		fmt.Println("work started", id, task)
		time.Sleep(1 * time.Second)
		fmt.Println("work finished", id, task)
		result <- task * 5
	}
}

func main() {

	const nooftasks = 10

	tasks := make(chan int, nooftasks)
	result := make(chan int, nooftasks)

	for work := 1; work <= 4; work++ {
		go worker(work, tasks, result)

	}

	for task := 1; task <= nooftasks; task++ {
		tasks <- task

	}
	close(tasks)

	for res := 1; res <= nooftasks; res++ {
		<-result

	}

}
