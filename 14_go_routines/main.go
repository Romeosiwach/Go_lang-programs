package main

// Go Routines ->
// 1. Go routine is a function or a method which execute independently and simultaneously in connection with any other goroutines present in the program //
// 2. Go Routine -> Any congurently executing activity in golanguage is called go routine //

// Congurency ->
// Congurency is the process of execution of multiple activities simultaneously together but one after another eg. Surfing, AC, Rice //

// Parallism ->
// Parallism is the process of execution of multiple activities simultaneously in parallel manner but executed by different users or cores //

import (
	"fmt"
	"sync"
)

// func main() {
// 	go fmt.Println("Romeo")
// 	go fmt.Println("Siwach")
// 	go fmt.Println("Saksham")

// 	// time.Sleep(1 * time.Second)
// 	fmt.Println("Vikas")
// 	fmt.Println("Parashar")

// }

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	user()

	fmt.Println("Aman")
	fmt.Println("Vikas")
	wg.Wait()
}
func user() {
	go fmt.Println("Romeo")
	go fmt.Println("Siwach")
	go fmt.Println("Ravi")
	wg.Done()
}
