package main

// FUNCTIONS ->
// Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code which ultimately saves the excessive use of memory, acts as a time saver and more importantly, provides better readability of the code. So basically, a function is a collection of statements that perform some specific task and return the result to the caller. A function can also perform some specific task without returning anything.//
//Function Expression -> func function_name(Parameter-list)(Return_type){function body.....}

import "fmt"

func main() {
	foo()

	bar("James")
	y := zoo("Saksham")
	fmt.Println(y)

	x, y := mouse("Romi", "Vikas")
	fmt.Println(x)
	fmt.Println(y)
}
func foo() {
	fmt.Println("hello from foo")
}
func bar(s string) {
	fmt.Println("hello,", s)
}
func zoo(x string) string {
	return fmt.Sprint("hello,", x)
}
func mouse(fn string, ln string) (string, string) {
	a := fmt.Sprint(fn, " says hello")
	b := fmt.Sprint(ln, " says hello")

	return a, b
}

//*********************************** Program for VARIADIC PARAMETERS in FUNCTIONS ****************************************//

// func main() {
// 	foo(2, 3, 4, 5, 6, 7, 8, 9)

// }
// func foo(x ...int) {
// 	sum := 0
// 	for i, v := range x {
// 		sum += v
// 		fmt.Println("At index position :", i, "We add the value :", v, "Total we get :", sum)
// 	}
// }

//********************************************Program for VARIADIC PARAMETERS in FUNCTIONS*********************************//

// func main() {
// 	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
// 	fmt.Println("total is", x)
// }
// func sum(x ...int) int {
// 	sum := 0
// 	for i, v := range x {
// 		sum += v
// 		fmt.Println("At index position :", i, "We add the value :", v, "Total we get :", sum)

// 	}
// 	fmt.Println("total is now", sum)
// 	return sum
// }
//********************************************Program for METHODS in FUNCTIONS*********************************//

// type person struct {
// 	first string
// 	last  string
// }
// type secretagent struct {
// 	person
// 	ltk bool
// }

// func main() {
// 	sa1 := secretagent{
// 		person: person{
// 			first: "James",
// 			last:  " Bond",
// 		},
// 		ltk: true,
// 	}
// 	sa2 := secretagent{
// 		person: person{
// 			first: "Miss",
// 			last:  " Moneypenny",
// 		},
// 		ltk: true,
// 	}
// 	fmt.Println(sa1)
// 	fmt.Println(sa2)
// 	sa1.speak()
// 	sa2.speak()
// }

// // func (r receiver) identifiers(parameters)(return(s)) {<code>} //

// func (s secretagent) speak() {
// 	fmt.Println("I am", s.first, s.last)

// }
