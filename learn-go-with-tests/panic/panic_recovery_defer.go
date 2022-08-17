package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		// get return of panic
		r := recover()
		if r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		panicNumber := fmt.Sprintf("%v", i)

		fmt.Println("Panicking!", panicNumber)
		// log.Fatalln(panicNumber)  // stop when call fatal , defer won't be executed
		panic(panicNumber)
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
