/*
Problem statement : fix race condition issue
in exc 3 using mutex
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	const numGoRoutines = 100
	wg.Add(numGoRoutines)
	for i := 0; i < numGoRoutines; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			mu.Unlock()
			wg.Done()
			fmt.Println("Number of go routines", runtime.NumGoroutine())
			fmt.Println("count value", count)
		}()
	}
	wg.Wait()
}
