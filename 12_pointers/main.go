package main

import (
	"fmt"
	"math"
)

// POINTERS -> Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable.
// Pointers in Golang is also termed as the special variables.
// The variables are used to store some data at a particular memory address in the system.
// The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF

//********************************************* POINTERS *****************************************************//

// func main() {
// 	x := 0
// 	fmt.Println("4", x)
// 	fmt.Println("4", &x)

// 	foo(&x)
// 	fmt.Println("3", x)
// 	fmt.Println("3", &x)

// }

// func foo(y *int) {

// 	fmt.Println("2", y)
// 	fmt.Println("2", &y)
// 	fmt.Println("2", *y)
// 	*y = 42
// 	fmt.Println("1", y)
// 	fmt.Println("1", &y)
// 	fmt.Println("1", *y)
// }

// ******************************************** METHOD SETS ***************************************************//

func main() {
	r := circle{
		radius: 5,
	}
	info(r)
}

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}
func info(s shape) {
	fmt.Println(s.area())
}
