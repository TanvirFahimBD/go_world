package main

import "fmt"

func update(a *int) {
	fmt.Println(a)
	*a = *a + 10
}

func main() {

	var x int
	var y *int

	fmt.Println("x value is ", x)
	fmt.Println("x memory address is ", &x)

	fmt.Println("y value is ", y)
	fmt.Println("y memory address is ", &y)

	x = 10
	y = &x

	fmt.Println("x value is ", x)
	fmt.Println("y referencing address is ", y)

	fmt.Println("y dereferencing value is ", *y)

	update(&x)

	fmt.Println(x)
}
