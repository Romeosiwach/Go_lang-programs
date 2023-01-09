package main

import "fmt"

// INTERFACE ->
// The interface is a collection of methods as well as it is a custom type.//

// POLYMORPHISM ->
// Polymorphism is the ability of a message to be displayed in more than one form.
// Polymorphism is considered as one of the important features of Object-Oriented Programming and can be achieved during either at runtime or compile time.//

//************************************Program for FUNCTIONS Interfaces and Polymorphism*******************************************//

//\\\\\\\\\\\\\\\\ TYPE 1\\\\\\\\\\\\\\\\\//
type person struct {
	first string
	last  string
}

//\\\\\\\\\\\\\\\\ TYPE 2\\\\\\\\\\\\\\\\\//
type secretagent struct {
	person
	ltk bool
}

//\\\\\\\\\\\\\\\\\ TYPE 3\\\\\\\\\\\\\\\\\//
type human interface {
	speak()
}

//\\\\\\\\\\\\\\\\ TYPE 4 \\\\\\\\\\\\\\\\//
type hotdog int

/////////////////// FUNCTION 1 ///////////////////////
func (s secretagent) speak() {
	fmt.Println("I am", s.first, s.last)
}

////////////////// FUNCTION 2 ///////////////////////
func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

////////////////// FUNCTION 3 ///////////////////////
func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

///////////////// FUNCTION 4 [ASSERTION] ///////////////////////
func car(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed in car", h.(person).first)
	case secretagent:
		fmt.Println("I was passed into car", h.(secretagent).first)

	}

}

///////////////// FUNCTION MAIN ////////////////////
func main() {
	sa1 := secretagent{
		person: person{
			first: "Romeo",
			last:  "Siwach",
		},
		ltk: true,
	}
	sa2 := secretagent{
		person: person{
			first: "Deepika",
			last:  "Siwach",
		},
		ltk: true,
	}
	p1 := person{
		first: "Saksham",
		last:  "Siwach",
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)
	sa1.speak()
	sa2.speak()
	p1.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
	car(sa1)
	car(sa2)
	car(p1)

	// CONVERSION //

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
