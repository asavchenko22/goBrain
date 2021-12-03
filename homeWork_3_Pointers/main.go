package main

import "fmt"

func main() {
	fmt.Println("_________first task__________")
	fmt.Println()

	var a *int
	var b = 55
	a = &b
	fmt.Println(*a)
	*a = 66
	fmt.Println(b)

	fmt.Println("_________second task_________")
	fmt.Println()

	var r *float64
	perimeter := 35.0
	pi := 3.14

	if perimeter == 35.0 {
		perimeter = (perimeter / 2.0 * pi)
		r = &perimeter
		fmt.Println(*r)
	}
}
