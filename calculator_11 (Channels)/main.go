package main

import "fmt"

func main() {
	x := 200
	y := 300
	fmt.Println(x)
	fmt.Println(y)
	z := make(chan int)

	go func() {

		for i := 0; i < y; i++ {
			fmt.Println(i)
			z <- i

		}
		close(z)
	}()
	for v := range z {
		fmt.Println(v)
	}

}
