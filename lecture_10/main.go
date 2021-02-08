package main

import "fmt"

func main() {

 x:= make(map[string]string)

 x["name"] = "Tanvir"
 x["age"] = "22"

 delete(x, "age") 

 fmt.Println(x)

}
package main

import "fmt"

func main() {

//  var students [3]string
//  students[0] = "Tanvir"
//  students[1] = "Hossain"
//  students[2] = "Fahim"

//  fmt.Println(students)
//  fmt.Println(len(students))
//  fmt.Println(students[0])
//  fmt.Println(students[1])
//  fmt.Println(students[2])


// students := [3]string{"Tanvir","Hossain","Fahim"}
// fmt.Println(students)

students := [...]string{"Tanvir","Hossain","Fahim","learning","Golang"}
 fmt.Println(students)


}
