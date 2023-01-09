package main

// Channels ->
// Channels are the methods through which one Goroutine communicate with another Goroutine and this communnication is lockfree //
// Two type of channels ->
// Buffered - Allows a limited number of send(chan<-) value with a receiver(<-chan) of those values //
// Unbuffered - Do not allow any send(chan<-) value without any receiver(<-chan) of those values //

import (
	"fmt"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++ Program for directional CHANNELS +++++++++++++++++++++++++++++++++++++ //
// func main() {
// 	{
// 		c := make(chan int, 2)
// 		c <- 42
// 		c <- 43
// 		fmt.Println(<-c)
// 		fmt.Println(<-c)
// 		fmt.Println("----")
// 		fmt.Printf("%T\n", c)
// 	}

// 	{
// 		c := make(chan int)
// 		cr := make(<-chan int) // receive //
// 		cs := make(chan<- int) // send //

// 		fmt.Printf("c\t:%T\n", c)
// 		fmt.Printf("cr\t:%T\n", cr)
// 		fmt.Printf("cs\t:%T\n", cs)

// 		// general to specific converts //
// 		fmt.Println("-----------")
// 		fmt.Printf("c\t:%T\n", (<-chan int)(c))
// 		fmt.Printf("c\t:%T\n", (chan<- int)(c))

// 	}

// }

//+++++++++++++++++++++++++++++++++++++++++ Program for Using CHANNELS +++++++++++++++++++++++++++++++++++++++++++//

// func main() {

// 	c := make(chan int)

// 	// SEND //
// 	go foo(c)

// 	// RECEIVE //
// 	bar(c)

// 	fmt.Println("about to exit")
// }

// // SEND //
// func foo(c chan<- int) {
// 	c <- 42
// }

// // RECEIVE //
// func bar(c <-chan int) {
// 	fmt.Println(<-c)
// }

//+++++++++++++++++++++++++++++++++++++++++ Program for Using RANGE over the CHANNELS +++++++++++++++++++++++++++++++++++++++++++//

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("abt to ext")
}

//+++++++++++++++++++++++++++++++++++++++++ Program for Using SELECTION over the CHANNELS +++++++++++++++++++++++++++++++++++++++++++//
// func main() {
// 	eve := make(chan int)
// 	odd := make(chan int)
// 	quit := make(chan int)

// 	// SEND //
// 	go send(eve, odd, quit)
// 	// RECEIVE //
// 	receive(eve, odd, quit)

// 	fmt.Println("About to quit")
// }

// func send(e, o, q chan<- int) {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			e <- i
// 		} else {
// 			o <- i
// 		}
// 	}
// 	q <- 0
// }
// func receive(e, o, q <-chan int) {
// 	for {
// 		select {
// 		case v := <-e:
// 			fmt.Println("From the even channel :", v)
// 		case v := <-o:
// 			fmt.Println("From the odd channel :", v)
// 		case v := <-q:
// 			fmt.Println("From the quit channel :", v)
// 			return
// 		}
// 	}
// }
//+++++++++++++++++++++++++++++++++++++++++ Program for Using COMMA OK over the CHANNELS +++++++++++++++++++++++++++++++++++++++++++//

//	func main() {
//		c := make(chan int)
//		go func() {
//			c <- 42
//			close(c)
//		}()
//		v,ok := <-c
//		fmt.Println(v, ok)
//		fmt.Println(<-c)
//	}
//
// +++++++++++++++++++++++++++++++++++++++++ Program for Using FAN IN in the CHANNELS +++++++++++++++++++++++++++++++++++++++++++//
// func main() {
// 	c := fanIn(boring("Joe"), boring("Ann"))
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-c)

// 	}
// 	fmt.Println("You are both boring; I'm leaving.")
// }

// Func boring //
// func boring(msg string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s%d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()
// 	return c
// }

// // Func fanIn //
//
//	func fanIn(input1, input2 <-chan string) <-chan string {
//		c := make(chan string)
//		go func() {
//			for {
//				c <- <-input1
//			}
//		}()
//		go func() {
//			for {
//				c <- <-input2
//			}
//		}()
//		return c
//	}
//
// ====================== FAN OUT Example ============================= //
// func main() {
// 	c1 := make(chan int)
// 	c2 := make(chan int)

// 	go populate(c1)
// 	go fanOutIn(c1, c2)

// 	for v := range c2 {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("About to Exit")
// }
// func populate(c chan int) {
// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}
// 	close(c)
// }
// func fanOutIn(c1, c2 chan int) {
// 	var wg sync.WaitGroup
// 	for v := range c1 {
// 		wg.Add(1)
// 		go func(v2 int) {
// 			c2 <- timeConsumingWork(v2)
// 			wg.Done()
// 		}(v)
// 	}
// 	wg.Wait()
// 	close(c2)
// }
// func timeConsumingWork(n int) int {
// 	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
// 	return n + rand.Intn(1000)
// }
