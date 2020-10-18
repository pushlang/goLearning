package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	/*
		n, k, l := 0, 0, 0
		   	for {
		   		n++
		   		select {
		   		case c <- rand.Int():
		   			k++
		   			//опциональный код здесь
		   		default:
		   			l++
		   			//тут можно ничего не писать, чтобы данные молча отбрасывались
		   			fmt.Println("выброшено")
		   		}
		   		time.Sleep(time.Millisecond * 5)
		   		if n==1000 {
		   			break
		   		}

		   		fmt.Printf("\n%d %d %d %f\n", n, k, l, (float32(l)/float32(n))*100)
		   	} */

	for {
		select {
		case c <- rand.Int():
		case <-time.After(time.Millisecond * 10):
			fmt.Println("тайм-аут")
		}
		time.Sleep(time.Millisecond * 1)
	}

}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("обработчик %d получил %d\n", w.id, data)
		time.Sleep(time.Millisecond * 50)
	}
}
