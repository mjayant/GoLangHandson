package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := make(chan int)

	go func() {
		c <- 56
		close(c)
		wg.Done()
	}()
	v, ok := <-c
	fmt.Println("value of v and ok", v, ok)
	v, ok = <-c
	fmt.Println("value of v and ok", v, ok)
	wg.Wait()
}
