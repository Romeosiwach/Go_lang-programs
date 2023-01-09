package main

// SWITCH STATEMENT ->
// The switch Statement Use the switch statement to select one of many code blocks to be executed called cases.And which case is found true then it is printed out.//
// Otherwise default case is printed out. //

import "fmt"

func main() {
	// EXAMPLE 1 //

	// var fruit string = "apple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("Its colour is red,its an apple")
	// case "banana":
	// 	fmt.Println("Its colour is yellow,its a banana")
	// case "watermelon":
	// 	fmt.Println("Its colour is green, its a watermelon")
	// default:
	// 	fmt.Println("Not a fruit")
	// }

	// EXAMPLE 2 //

	// var a int

	// fmt.Println("Enter number")
	// fmt.Scanf("%d\n", &a)

	// switch {
	// case a > 90:
	// 	fmt.Println("It is a gold medal")
	// case a > 70:
	// 	fmt.Println("It is a silver medal")
	// case a > 50:
	// 	fmt.Println("It is an bronze medal")
	// default:
	// 	fmt.Println("Fail")

	// }

	// var user string

	// fmt.Println("Enter name")
	// fmt.Scanf("%s\n", &user)

	// switch user {

	// case "romeo":
	// 	fmt.Println("Valid user")
	// case "vinay":
	// 	fmt.Println("Invalid user")
	// default:
	// 	fmt.Println("Not sure")
	// }

	var a int

	fmt.Println("Enter number")
	fmt.Scanf("%d\n", &a)

	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	}
}

// ******************************************** Switch Statement ***************************************************//

func main() {
	switch 2 == 0 {
	case 2 == 2:
		fmt.Println("Expression is true1")
	case 2 < 2:
		fmt.Println("Expression is false2")
	case 2 > 2:
		fmt.Println("The statement is false3")
	case 2 == 0:
		fmt.Println("Statement is false4")
	default:
		fmt.Println("Invalid")
	}

}
