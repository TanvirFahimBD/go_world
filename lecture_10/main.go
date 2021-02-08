package main

import "fmt"

func main() {

 x:= make(map[string]string)

 x["name"] = "Tanvir"
 x["age"] = "22"

 delete(x, "age") 

 fmt.Println(x)

}
