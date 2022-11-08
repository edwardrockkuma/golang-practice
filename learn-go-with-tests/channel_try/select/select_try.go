package main

import (
	"fmt"
	"time"
)

func doTask() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "finish"
	}()

	for {
		select {
		case (<-ch):
			fmt.Println("main is over")
			return
		default:
			fmt.Println("wait...")
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func main() {
	doTask()
}
