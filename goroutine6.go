package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data = make(map[string]string)
	rw   = &sync.RWMutex{}
)

func read(key string) string {
	rw.RLock()
	defer rw.RUnlock()
	return data[key]
}

func write(key, value string) {
	rw.Lock()
	defer rw.Unlock()
	data[key] = value
}

func main() {

	go read("foo")
	go read("bar")

	go write("baz", "qux")
	go read("baz")

	time.Sleep(time.Second)
	fmt.Println(data, rw)
}
