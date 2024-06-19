package main

import (
	"fmt"
	"sync"
)

var a = 1

func Sum(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	a = a + 1
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Sum(&wg, &mu)
	}
	wg.Wait()
	fmt.Println(a)
}
