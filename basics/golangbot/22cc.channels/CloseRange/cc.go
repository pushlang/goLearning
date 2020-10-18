//The for range form of the for loop can be used to receive values from a channel until it is closed.

package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
		ch := make(chan int)
	   	go producer(ch)
	   	for v := range ch {
	   		fmt.Println("Received ", v)
	   	}

	//Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed. ok=true if the value was received by a successful send operation. ok=false - reading from a closed channel. The value read from a closed channel will be the zero value of the channel's type.

	ch2 := make(chan int)
	go producer(ch2)
	for {
		v, ok := <-ch2
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

}

