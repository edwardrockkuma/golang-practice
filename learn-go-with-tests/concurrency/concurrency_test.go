package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebChecker(url string) bool {
	if url == "http://mock-host" {
		return false
	}
	return true
}

func TestCheckWebSite(t *testing.T) {
	webSite := []string{
		"http://google.com",
		"http://mock-host",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://mock-host":  false,
	}

	got := CheckWebsite(mockWebChecker, webSite)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func slowStubWebChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	//
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebChecker, urls)
	}
}
