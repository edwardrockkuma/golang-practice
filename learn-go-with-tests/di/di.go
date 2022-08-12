package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// io.Writer implement this interface
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

//var GreetMsg string = "hi~%s"

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hi~%s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "http")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler)))
}
