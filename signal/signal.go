package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 3)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	
	go func() {
		time.Sleep(1*time.Second)
		signal.Stop(c)
		fmt.Println("from now signal will not be interrupted!")
	}()
	
	signal.Notify(c)
	// signal.Notify(c, os.Interrupt)
	

	// Block until any signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}