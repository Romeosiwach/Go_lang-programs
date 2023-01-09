package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to Romeo's calculator")

	var firstnum, secondnum float32
	var options int
	var res int

	fmt.Println("Enter first num", firstnum, res)
	fmt.Scanf("%d\n", &firstnum)
	fmt.Println("Enter second num", secondnum, res)
	fmt.Scanf("%d\n", &secondnum)

	switch options {
	case 1:
		addition()
		fmt.Println("res\n : ", &firstnum, &secondnum, &res)

	case 2:
		substraction()
		fmt.Println("res\n : ", &firstnum, &secondnum, &res)

	case 3:
		multiplication()
		fmt.Println("res\n : ", &firstnum, &secondnum, &res)

	case 4:
		division()
		fmt.Println("res\n : ", &firstnum, &secondnum, &res)
	default:
		os.Exit(0)

	}
}

var a, b int

func addition() {
	addition := a + b
	fmt.Println(addition)
}
func substraction() {
	substraction := a - b
	fmt.Println(substraction)
}
func multiplication() {
	multiplication := a * b
	fmt.Println(multiplication)
}
func division() {
	division := a / b
	fmt.Println(division)
}
