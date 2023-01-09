package main

// FOR LOOP ->
// A for loop is a repetition control structure. It allows you to write a loop that needs to execute a specific number of times.
// Syntax- for [Init condition; Incremental ; Range]
// If you want to run the same code over and over again, each time with a different value. Each execution of a loop is called an ITERATION.

import "fmt"

//********************************************** For Loop **********************************************************//
func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

}
