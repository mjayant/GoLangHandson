package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
	wg.Wait()
}

func gen(q chan int) chan int {
	c := make(chan int)
	wg.Add(1)
	go func() {
		for i := 1; i < 101; i++ {
			c <- i
		}
		close(c)
		q <- 1
		wg.Done()
	}()
	return c
}

func receive(a chan int, b chan int) {
	for {
		select {
		case v := <-a:
			fmt.Println("Extract values from channel", v)

		case <-b:
			return
		}
	}
}
