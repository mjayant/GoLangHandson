package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var chanel_list []chan int
	wg.Add(10)
	for i := 0; i < 10; i++ {
		chanel_list = append(chanel_list, make(chan int))
	}

	for _, channel := range chanel_list {
		go produce(channel)
	}
	out := make(chan int)
	go FanIn(chanel_list, out)
	for val := range out {
		fmt.Println("Extracting value from out channel", val)
	}
	wg.Wait()
	fmt.Println("About to exit")
}

func produce(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	wg.Done()
}

func FanIn(a1 []chan int, out chan int) {
	var fanInWg sync.WaitGroup
	fanInWg.Add(len(a1))
	for _, channel := range a1 {
		go func(c chan int) {
			for val := range c {
				out <- val
			}
			fanInWg.Done()
		}(channel)

	}
	fanInWg.Wait()
	close(out)

}
