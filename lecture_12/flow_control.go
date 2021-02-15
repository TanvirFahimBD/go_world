package main

import (
	"fmt"
)

func main() {
	/*
	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d", &age)
	*/

	/*
	if boolean_expression{
		logic or statement Here
	}
	*/


	/*
	if age < 3 {	//0 to 2
		fmt.Println("infant")
	}else if age>2 && age<13{	//2 to 12
		fmt.Println("Children")

	}else if  age>12 && age<=19{	//13 to 19
		fmt.Println("Teenage")
	}else{				//20+
		fmt.Println("Adult")
	}	
	*/
	
	/*
	//fixed value
	switch age {
	case 2:
		fmt.Println("infant")
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Children")
	case 13,14,15,16,17,18,19:
		fmt.Println("Teenage")
	default:
		fmt.Println("Adult")
	}
	*/
	/*
	switch age {
	case 2:
		fmt.Println("infant")
		fallthrough
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Children")
	case 13,14,15,16,17,18,19:
		fmt.Println("Teenage")
	default:
		fmt.Println("Adult")
	}
	*/
	
	/*
	//for loop
	 //1,2,3,4,5,6,7,8,9
	for i:=1; i<=9; i++{
	fmt.Println(i)
	}
	*/
	
	/*
	students := []string{"Tanvir","Hossain","Fahim"}
	
	for i, std := range students {
		fmt.Println(i, std)
	}	
	*/

	i:=0
	for i<10 {
		fmt.Println(i, "Hello")		
	i++
	}

	
}