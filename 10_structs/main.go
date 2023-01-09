package main

// Struct ->
// A struct is a collection of data types and store them in the form of KEY and VALUES //
//

import "fmt"

//**************************************** Program for struct example no 0 *******************************************//

type user struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	romeo := user{"Romeo", "romeosiwach@gmail.com", 16, true}
	fmt.Println(romeo)
	fmt.Printf("Romeo's details are : %+v\n", romeo)
	fmt.Printf("Name is %v and email is %v", romeo.Name, romeo.Email)
}

// **************************************** Program for STRUCT example no1********************************************//
// import "fmt"

// func main() {

// 	type User struct {
// 		Name   string
// 		Email  string
// 		Status bool
// 		Age    int
// 	}
// 	Romeo := User{"Romeo", "Romeosiwach@gmail.com", true, 31}
// 	fmt.Println(Romeo)
// 	fmt.Printf("%+v\n", Romeo)
// 	fmt.Printf("%v\n", Romeo)
// 	fmt.Printf("%T\n", Romeo)

// }

// **************************************** Program for STRUCT example no2********************************************//

// func main() {
// 	type User struct {
// 		First string
// 		Last  string
// 	}
// 	Romeo := User{"romeo", "siwach"}
// 	fmt.Println(Romeo)
// }

// **************************************** Program for STRUCT example no3********************************************//

// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }
// type secretagent struct {
// 	person
// 	Ltk bool
// }

// func main() {
// 	sa1 := secretagent{

// 		person: person{
// 			First: "James",
// 			Last:  "Bond",
// 			Age:   27,
// 		},
// 		Ltk: true,
// 	}
// 	p2 := person{
// 		First: "Miss",
// 		Last:  "Moneypenny",
// 		Age:   35,
// 	}
// 	fmt.Println(sa1.person, sa1.Ltk)
// 	fmt.Println(p2.Age, p2.First, p2.Last)

// }
//************************************************************************************************************************************************//

// STRUCT ->
// A structure or struct in Golang is a user-defined type that allows to combine different types into a single type.//
// Golang program to show how to
// declare and define the struct

// Defining a struct type
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := Address{"Akshay", "Dehradun", 3623572}

	fmt.Println("Address1: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{Name: "Anikaa", city: "Ballia",
		Pincode: 277001}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)
}
