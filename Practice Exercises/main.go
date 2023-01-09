package main

import (
	"fmt"
)

//********************************************************* Functions *************************************************************//
// func main() {
// 	foo()
// 	bar("Monty")

// 	a := zoo("Ziyan")
// 	fmt.Println(a)

// 	b, c := moo("Rohit", "Mohit")
// 	fmt.Println(b)
// 	fmt.Println(c)

// }

// func foo() {
// 	fmt.Println("Hello this is foo")
// }
// func bar(s string) {
// 	fmt.Println("Hello this is ", s)
// }
// func zoo(x string) string {
// 	return fmt.Sprint("Welcome ,", x)
// }
// func moo(fn string, ln string) (string, string) {
// 	y := fmt.Sprintln("My name is,", fn)
// 	z := fmt.Sprintln("My name is,", ln)

// 	return y, z

// }
// ****************************************************** CHANNELS *****************************************************************//

// func main() {
// 	c := make(chan int)
// 	go foo(c)
// 	bar(c)
// 	fmt.Println("Exit")
// }

// func foo(c chan<- int) {
// 	c <- 42
// }
// func bar(c <-chan int) {
// 	fmt.Println(<-c)
// }

//Exercise 2//

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 20; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()

// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

// ******************************************************** GO ROUTINES *******************************************************//

// func main() {
// 	go fmt.Println("Romeo")
// 	go fmt.Println("Rohit")

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Monty")
// 	fmt.Println("Mohan")
// }

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(3)
// 	foo()
// 	fmt.Println("Monty")
// 	fmt.Println("Jonty")
// 	wg.Wait()
// }
// func foo() {
// 	go fmt.Println("Saksham")
// 	go fmt.Println("Surender")
// 	go fmt.Println("Sandra")
// 	wg.Done()
// }

//************************************************************* Struct ****************************************************//

type Address struct {
	Name    string
	City    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"Akshay", "Delhi", 110001}
	fmt.Println("Address1", a1)

	a2 := Address{"Rocky", "Mumbai", 220001}
	fmt.Println("Address2", a2)
}

//*****************************************Pointers*******************************************************//

func main(){
	var a int = 200
	
}
