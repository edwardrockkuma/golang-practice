package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	for i := 0; i <= 20; i++ {
		intChan <- i
		fmt.Printf("write data = %v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("read data = %v\n", v)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	intChan := make(chan int)
	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
}
