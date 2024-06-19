package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) { // workers will receive work on the ch1 channel and send the corresponding results on ch2
	for j := range jobs {
		fmt.Println(id, j)
		time.Sleep(time.Second)
		fmt.Println(id, j)
		results <- j + "done"
	}
}

func main() {

	const numJobs = 5

	jobsList := []string{"job1", "job2", "job3", "job4", "job5"}

	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for _, j := range jobsList {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
