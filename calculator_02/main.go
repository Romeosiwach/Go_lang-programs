package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Enter First Number")
	fmt.Scanf("%d\n", &a)

	fmt.Println("Enter Second Number")
	fmt.Scanf("%d\n", &b)

	Total := a + b

	fmt.Println("Result :", Total)

}
