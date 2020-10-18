//BUFFERED CHANNELS - created by passing an additional capacity parameter (size of the buffer) to the make function. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}

	{
		//CAPICITY - number of values that the channel can hold (value of the buffered channel). LENGTH - number of elements currently queued in it.
		ch := make(chan string, 3)
		ch <- "naveen"
		ch <- "paul"
		fmt.Println("capacity is", cap(ch))
		fmt.Println("length is", len(ch))
		fmt.Println("read value", <-ch)
		fmt.Println("new length is", len(ch))
	}
}
