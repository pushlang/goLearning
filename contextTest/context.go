package main

import (
    "log"
    "time"
    "context"
)

func main() {
	mainProcess() 
}

func mainProcess() {
    ctx, cancel := context.WithCancel(context.Background())
    
	go func() {
		log.Println("a waiting before processes will be cancelled!")
		time.Sleep(3 * time.Second)
		cancel()
	}()

	log.Println("start of subProcess!")
    err := subProcess(ctx, 1)

	log.Println(err, "by the mainProcess!")
}

func subProcess(ctx context.Context, id int) error {
	cx, cancel := context.WithCancel(ctx)
	defer cancel()
	
 	c := make(chan error)
	go func() {
		c<- longTimedRequest(cx)
	}()
	
	select {
	case <-c:
		log.Println("subProcess finished!")
	case <-ctx.Done():
		log.Println("subProcess cancelled! With err:", ctx.Err())
		return ctx.Err()
	}

	return nil
}

func longTimedRequest(ctx context.Context) error {
	cx, cancel := context.WithCancel(ctx)
	defer cancel()

	tick := time.Tick(1 * time.Second)
	for i:=0; i<10; i++ {
		select {
		case <-tick:
			log.Println("longTimedRequest tick!")
		case <-cx.Done():
			log.Println("longTimedRequest cancelled! With err:", ctx.Err())
			return cx.Err()
		}
	}
	
	log.Println("longTimedRequest finished!")
	return nil
}