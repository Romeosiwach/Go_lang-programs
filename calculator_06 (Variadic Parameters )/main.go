package main

import "fmt"

//********************************************Program for VARIADIC PARAMETERS in FUNCTIONS*********************************//
func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("total is", x)
}
func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("At index position :", i, "We add the value :", v, "Total we get :", sum)

	}
	fmt.Println("total is now", sum)
	return sum
}

// ***************************************** Type 2 ************************************************* //

func main() {
	foo(2, 3, 4, 5, 6, 7, 8, 9)
}
func foo(x ...int) {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("At index position :", i, "We add the value :", v, "Total we get :", sum)
	}
}
