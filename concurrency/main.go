/*
Problem statement :
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go foo()
	go bar()
	fmt.Println("Number of goroutines are active \t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Number of CPUs are active \t", runtime.NumCPU())
}

func foo() {
	for i := 0; i < 101; i++ {
		fmt.Println("---------Inside foo-----", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 101; i++ {
		fmt.Println("---------Inside bar-----", i)
	}
	wg.Done()
}
