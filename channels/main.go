package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		c <- 45
		wg.Done()
	}()
	go func() {
		fmt.Println("Value in channel", <-c)
		wg.Done()
	}()
	fmt.Println("Exit from main")
	wg.Wait()
}
