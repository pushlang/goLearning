//WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing.The Wait() method blocks the Goroutine in which it's called until the counter becomes zero. The way to decrement the counter is by calling Done() method on the WaitGroup.

package main

import (  
    "fmt"
    "sync"
    "time"
)

func process(i int, wg *sync.WaitGroup) {  
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done()
}

func main() {  
    no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
		go process(i, &wg)
		//It is important to pass the address of wg in line no. 21. If the address is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")
}