package concurrency

import (
	"fmt"
	"time"
)

type WebChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	time.Sleep(1 * time.Second)

	return results
}

func Print(t time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(t)
	}
}
