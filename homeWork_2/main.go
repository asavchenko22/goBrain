package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "104"
	num := 35
	num1, _ := strconv.Atoi(str)
	fmt.Println(num1)
	fmt.Println(strconv.Itoa(num))
}
