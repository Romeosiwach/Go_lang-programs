// **********************************Program for FUNCTIONS of type  FUNCTIONS RECURSION  ************************************//
package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	r := factorial(4)
	fmt.Println(r)
	s := loopfact(4)
	fmt.Println(s)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func loopfact(n int) int {
	total := 1

	for ; n > 0; n-- {
		total *= n
	}
	return total
}

// type 2 //

package main

import "fmt"

func main() {

	var factorialnum, factorial int
	factorial = 1

	fmt.Print("Enter any Number to find the Factorial = ")
	fmt.Scanln(&factorialnum)

	for i := 1; i <= 10; i++ {
		factorial = factorial * i
	}
	fmt.Println("The Factorial of ", 10, " = ", factorial)
}
