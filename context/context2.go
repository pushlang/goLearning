package main

import (
    "fmt"
    "net/http"
    "time"
    "context"
)

func main() {
	mainProcess() 
}

func mainProcess() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        time.Sleep(5 * time.Second)
        ctx.Done() <- struct{}{}
    }()
   
    res1, err:= subProcess(ctx, 1)
	//time.Delay(3 * time.Second)
	//res2, err:= go subProcess(ctx, 2)
    
	if err != nil {
        http.Error(w, "cancelled", http.StatusInternalError)
        return
    }
	
	//fmt.Println("res1:", res1, "res2:", res2)
}

func subProcess(ctx context.Context, id int) (int, error) {
 	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("Tick id#", id, countdown)
		case <-ctx.Done():
			fmt.Println("Tick id#", id, "stopped!")
			return nil, ctx.Err()
		}
	}
}