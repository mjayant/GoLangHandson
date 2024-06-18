/*
Problem statement
write a program that
○ launches 10 goroutines
■ each goroutine adds 10 numbers to a channel
○ pull the numbers off the channel and print them
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			var wg1 sync.WaitGroup
			wg1.Add(1)
			c1 := make(chan int)
			go merge(c1, c, &wg1)
			for i := 1; i < 11; i++ {
				c1 <- i
			}
			go func() {
				close(c1)
				wg1.Wait()
			}()
			wg.Done()
		}()

	}

	fmt.Println("Extracting value from common channel")
	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("About to exit")
	go func() {
		close(c)
		wg.Wait()
	}()
}

func merge(c1 chan int, c chan int, wg1 *sync.WaitGroup) {
	fmt.Println("Sending channels value to common channel")
	for val := range c1 {
		c <- val
	}
	wg1.Done()
}
