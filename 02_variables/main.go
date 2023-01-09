package main

// VARIABLE ->
// Variables are the storage types in which the values can be changed during the program execution according to need. "Var"//
// Examples ->
// var a int = 99
// var b = 100
// var (
// 	ab = 99
// 	bb = 89
// )
// var (
// 	Bb2 int     = 78
// 	cc  float32 = 7.7
// )

import (
	"fmt"
	//"os"
)

// EXAMPLE - 01 //
// func main() {

// 	var x int = 99
// 	var y = 100
// 	var (
// 		a         = 10
// 		b         = 20
// 		c float32 = 7.7
// 		d float32 = 32
// 	)

// 	Aksf := 0

// 	fmt.Println("Sum of two number", x+y)
// 	fmt.Println("Sum of two number", a+b)
// 	fmt.Println("Sum of two number", c+d)
// 	fmt.Println("The declared number is", Aksf)

// }

//ROMEO //

func main() {

	//CONDITION 1 //
	// 	var (
	// 		a = 10
	// 		b = 20
	// 	)

	// 	Total := a + b
	// 	fmt.Println("Total :", Total)
	// }

	// 	//CONDITION 2 //
	// 	var a, b string

	// 	fmt.Println("Choose the options")
	// 	fmt.Println("red")
	// 	fmt.Println("yellow")
	// 	fmt.Scanf("%d\n", &a)
	// 	fmt.Scanf("%d\n", &b)

	// 	colour := a

	// 	if colour = "red" {
	// 		fmt.Println("It is an Apple")
	// 	} else if colour = "yellow" {
	// 		fmt.Println("It is a Banana")

	// 	} else {
	// 		os.Exit(0)
	// 	}
	// }

	// 		//CONDITION 3 //
	colour := "green"
	taste := "sweet" // (declared and initialised)

	if colour == "red" && taste == "sweet" {
		fmt.Println(" It is an apple")
	} else {
		fmt.Println("exit")
	}

}
