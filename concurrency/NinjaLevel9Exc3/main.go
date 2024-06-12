/*
problem stetement :
● Using goroutines, create an incrementer program
○ have a variable to hold the incrementer value
○ launch a bunch of goroutines
■ each goroutine should
● read the incrementer value
○ store it in a new variable
● yield the processor with runtime.Gosched()
● increment the new variable
● write the value in the new variable back to the incrementer
variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int
var wg sync.WaitGroup

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := count
			v++
			count = v
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Number of GoRoutine", runtime.NumGoroutine())
		fmt.Println("Count value", count)
	}
	wg.Wait()
}
