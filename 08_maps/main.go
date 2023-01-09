package main

import "fmt"

// MAPS ->
// Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys. It is a reference to a hash table. //

//*********************************** Programs for making MAPS ************************************************//

// func main() {

// 	m := map[string]int{
// 		"romi": 30,
// 		"deep": 26,
// 	}
// 	fmt.Println(m)
// 	fmt.Println(m["romi"])
// 	fmt.Println(m["deep"])
// 	fmt.Println(m["robert"])

// 	v, ok := m["robert"] // Formula to check whether details present in our data or not //

// 	fmt.Println(v)
// 	fmt.Println(ok)
// }

//*********************************** Programs for making MAPS using IF statement ************************************************//

// func main() {

// 	m := map[string]int{
// 		"romi": 30,
// 		"deep": 26,
// 	}
// 	m["saksham"] = 4

// 	if v, ok := m["romi"]; ok {
// 		fmt.Println(v)
// 	}
// }

//*********************************** Programs for  MAPS Deletion ************************************************//

// func main() {
// 	m := map[string]int{
// 		"romi": 31,
// 		"deep": 26,
// 	}
// 	fmt.Println(m)
// 	delete(m, "romi")
// 	fmt.Println(m)
// 	fmt.Println(m["romi"])
// 	fmt.Println(m["deep"])

// 	if v, ok := m["deep"]; ok {
// 		fmt.Println("value :", v)
// 		delete(m, "romi")

// 	}
// 	fmt.Println(m)
// }

// ********************************************* Program for practice ************************************************//

func main() {
	languages := make(map[string]string)
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["JS"] = "JavaScript"

	fmt.Println(languages)

	for i, v := range languages {
		fmt.Printf("The term %v stands for %v\n", i, v)
	}

	for _, v := range languages {
		fmt.Printf("The term v stands for %v\n", v)
	}

}
