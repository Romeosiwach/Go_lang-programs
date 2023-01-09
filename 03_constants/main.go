package main

// CONSTANTS ->
// Constants are storage types but in these once the value of constant is defined it cannot be modified further //

import "fmt"

func main() {

	//CONDITION 1
	// const
	// 	a = 10
	// 	b = 20

	// Total := a + b
	// fmt.Println("Total :", Total)

	//CONDITION 2
	// const taste, colour = "sweet", "red"

	// if colour == "red" {
	// 	fmt.Println("It is an Apple")

	//CONDITION 3
	// colour := "blue"
	// taste := "sweet" // (declared and initialised)

	// if colour == "red" && taste == "sweet" {
	// 	fmt.Println(" It is an apple")
	// } else {
	// 	fmt.Println("exit")
	//}

	//CONDITION 4
	const a int = 4
	const b = 5
	const c = 7

	Total := a + b + c

	fmt.Println(Total)
	fmt.Scanf("%d\n", Total)
}
