package main

// Slice ->
// Slice is the collection of elements of same type but the size of slices can be resized unlike arrays //
//1. fruitList := []string{"Apple", "Mango", "Grapes"} //
//2. numberList := make([]int, 4) //
//   numberList[0] = 215 //

import "fmt"

func main() {
	// fruitList := []string{"Apple", "Mango", "Grapes"}
	// fmt.Println("Here is the fruit list :", fruitList)

	// fruitList = append(fruitList, "Guavava", "Watermelon")
	// fmt.Println("Here is the new fruit list :", fruitList)
	// fmt.Printf("Here is the new fruit list type : %T\n", fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println("Here is the new fruit list :", fruitList)

	numberList := make([]int, 4)

	numberList[0] = 215
	numberList[1] = 330
	numberList[2] = 445
	numberList[3] = 560
	fmt.Println("This in the numberlist :", numberList)

	numberList = append(numberList, 675, 790, 805)
	fmt.Println(numberList)
}

// ******************************* Program for using SLICE and printing its RANGE,LENGTH,INDEX-VALUE *****************************************//
// import "fmt"

// func main() {
// 	x := []int{2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(x[0])
// 	fmt.Println(x[1])
// 	fmt.Println(x[2])
// 	fmt.Println(x[3])
// 	fmt.Println(x[4])
// 	fmt.Println(x[5])
// 	fmt.Println(x[6])

// 	for i, v := range x {
// 		fmt.Println(i, v)
// 	}
// }

// ************************************************* Program for Printing ARRAY *************************************************************//

// func main() {
// 	var x [5]int
// 	var Y [6]int

// 	fmt.Println(x)
// 	fmt.Println(Y)
// 	x[3] = 42

// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// }

// ********************************************Program for Printing SLICES *************************************************//

//x:= type {value}  // composite literals // [] => Slice //

// func main() {
// 	x := []int{2, 3, 4, 5, 6, 7}
// 	fmt.Println(x)
// }

//****************************************** Program for printing SLICE RANGE  *******************************************//

// func main() {
// 	x := []int{2, 3, 4, 5, 6, 7}

// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	fmt.Println(x[0])
// 	fmt.Println(x[1])
// 	fmt.Println(x[2])
// 	fmt.Println(x[3])

// 	for i, v := range x {
// 		fmt.Println(i, v)                 // i => index, v => value  //
// 	}
// }

//********************************************* Program for printing SLICING A SLICE **************************************//

// func main() {
// x := []int{2, 3, 4, 5, 6, 7}
// fmt.Println(x)
// fmt.Println(x[3])
// fmt.Println(x[1:])
// fmt.Println(x[1:3])

// for i, v := range x {
// 	fmt.Println(i, v)
//}
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i, x[i])
// 	}
// }

//******************************************** Program for using APPEND TO SLICE ******************************************** //

// func main() {
// 	x := []int{1, 2, 3, 4, 5, 6}

// 	fmt.Println(x)

// 	x = append(x, 77, 78, 79, 80) // Type 1 //
// 	fmt.Println(x)

// 	y := []int{10, 11, 12, 13, 14}

// 	x = append(x, y...)

// 	fmt.Println(x)
// }

//********************************** Program for SCLICE DELETING FROM A SLICE ********************************************//

// func main() {
// 	x := []int{2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(x)
// 	x = append(x[:2], x[4:]...)
// 	fmt.Println(x)
// }

//****************************************** Program for  SLICE MAKE ********************************************//
// func main() {

// 	x := []int{2, 3, 4, 5, 6, 7, 8, 9}

// 	x = make([]int, len(x), cap(x))
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x[3] = 42 // We can store new values at specific places //
// 	x[5] = 999
// 	fmt.Println(x)

// 	x = append(x, 5560) // We can store new vaues while adding new places in the slice//
// 	fmt.Println(x)
// }

//****************************** Program for multi-dimensional SLICES ********************************************//

// func main() {
// 	rs := []string{"romi", "siwach", "chocolate", "vanilla"}
// 	ds := []string{"deep", "siwach", "strawberry", "hazelnut"}

// 	x := [][]string{rs, ds}
// 	fmt.Println(rs)
// 	fmt.Println(ds)
// 	fmt.Println(x)
// }

//*******************************************************************************************************************//
