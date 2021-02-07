package main

import "fmt"

func main() {

 fmt.Println("Enter Your name and age: ")
 var name string
 var age int
 fmt.Scanf("%s %d", &name, &age)
 fmt.Printf("Your name %s and age is %d", name, age)
	
}
