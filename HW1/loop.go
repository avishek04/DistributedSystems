package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	numRoutines := 2
	wg.Add(numRoutines)

	for j := 0; j < 2; j++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
