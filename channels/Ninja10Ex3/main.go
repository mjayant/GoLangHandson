package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := gen()
	go reciver(c)
	wg.Wait()
}

func gen() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func reciver(c chan int) {
	for val := range c {
		fmt.Println(" Extract value from channel", val)
	}
	wg.Done()
}
