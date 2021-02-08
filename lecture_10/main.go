// package main

// func main() {

// 	//  var students [3]string

// 	//  students[0] = "Tanvir"
// 	//  students[1] = "Hossain"
// 	//  students[2] = "Fahim"

// 	//  x := students[0:2]
// 	//  fmt.Println(x)

// 	// var students [3]string

// 	// x := make([]string, 0)
// 	// fmt.Println(x)
// 	var students [3]string
// 	var fruits []string
// 	fruits = append(fruits, "Apple", "Mango")

// 	fmt.Println(fruits, len(fruits))

// 	fmt.Printf("%T\n", fruits)
// 	fmt.Printf("%T", students)

// }

//Identify slice code

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var students [3]string
	var fruits []string
	fruits = append(fruits, "Apple", "Mango")

	fmt.Println(fruits, len(fruits))

	fmt.Printf("%T\n", fruits)
	fmt.Printf("%T\n", students)

	a := reflect.TypeOf(students).Kind().String()
	b := reflect.TypeOf(fruits).Kind().String()
	fmt.Println(a)
	fmt.Println(b)

}
