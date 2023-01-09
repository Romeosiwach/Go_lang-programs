package main

import (
	"fmt"
)

func main() {

	// fmt.Println(runtime.GOOS)
	// fmt.Println(runtime.GOARCH)

	fmt.Println("Hello World !")

	zoo()
	loo()
}

func zoo() {
	foo := 100
	fmt.Println(foo)

}
func loo() {

	for i := 0; i < 10; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}

	}

}
