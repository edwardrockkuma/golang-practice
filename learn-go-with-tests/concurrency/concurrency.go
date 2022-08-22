package concurrency

import (
	"fmt"
	"time"
)

type WebChecker func(string) bool

func CheckWebsite(wc WebChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		result[url] = wc(url)
	}

	return result
}

func Print(t time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(t)
	}
}
