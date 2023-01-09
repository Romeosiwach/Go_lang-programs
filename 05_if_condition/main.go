package main

// IF Condition ->
// The if-else statement allows you to execute one block of code if the specified condition is evaluates to true and another block of code if it is evaluates to false.//

import "fmt"

func main() {

	// Example 1//

	// var a, b int

	// fmt.Println("Enter firt number")
	// fmt.Scanf("%d\n", &a)
	// fmt.Println("Enter second number")
	// fmt.Scanf("%d\n", &b)

	// if a >= b {
	// 	fmt.Println("Result :", a-b)

	// } else if b >= a {
	// 	fmt.Println("exit")
	// }

	// Example 2//

	var colour, taste string = "blue", "sweet"

	if colour == "red" {
		fmt.Println("It is an apple")

	} else if taste == "sour" {
		fmt.Println("It is an orange")

	} else {
		fmt.Println("Exit")
	}

}

// ************ Program for printing FOR{FOR} LOOP *************//
// func main() {
// 	for i := 0; i <= 5; i++ {
// 		for j := 0; j < 3; j++ {

// 			// fmt.Printf("%d\n%d\n%d\n", i, j, k)
// 			fmt.Printf("external most : %d\t intermediate : %d\n ", i, j)

// 		}
// 	}

// }

// ******** Program using BREAK and COTINUE statements *************//

// func main() {

// 	for x := 1; x < 100; x++ {

// 		if x > 20 {
// 			break
// 		}
// 		if x%2 != 0 {
// 			continue
// 		}
// 		fmt.Println(x)
// 	}
// }

//************ Program to print HEXADECIMAL VALUE And ASCII characters ************* //

// func main() {

// 	for x := 30; x < 50; x++ {

// 		fmt.Printf("%v\t%#U\n", x, x)
// 	}
// }

// ************* Program to print Conditional statement --- IF *****************//

// func main() {

// 	if true {
// 		fmt.Println(001)
// 	}
// 	if false {
// 		fmt.Println(002)
// 	}
// 	if !true {
// 		fmt.Println(003)
// 	}
// 	if !false {
// 		fmt.Println(004)
// 	}
// 	if 2 == 2 {
// 		fmt.Println(005)
// 	}
// 	if 2 < 2 {
// 		fmt.Println(006)
// 	}
// 	if 2 <= 2 {
// 		fmt.Println(007)
// 	}
// }

//*********** Program to print Initialisation in IF statement ***********//
// func main() {

// 	if x := 42; x == 2 {
// 		fmt.Println(001)
// 	}
// 	if x := 2; x == 2 {
// 		fmt.Println(002)
// 	}

// }

//*********** Program to print Initialisation in IF, else IF, else....statements ***********//

func main() {
	x := 11
	if x <= 10 {
		fmt.Println("Less then 10")
	} else if x >= 10 {
		fmt.Println("Greater then 10")
	} else if x == 10 {
		fmt.Println("Correct")
	} else {
		fmt.Println("INCORRECT")
	}
}
