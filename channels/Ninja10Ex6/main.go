/*
Problem statement :
write a program that
> puts 100 numbers to a channel
> pull the numbers off the channel and print them
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(2)
	go func() {
		for i := 1; i < 101; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}()
	reciver(c)
	fmt.Println("About to exit")
	wg.Wait()
}

func reciver(c chan int) {
	var wg1 sync.WaitGroup
	wg1.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("Extract value from channel", <-c)
			wg1.Done()
		}()
	}
	wg1.Wait()
	wg.Done()
}
