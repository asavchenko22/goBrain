package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("________the first part of assignment________")
	fmt.Println()

	str := "104"
	num := 35
	num1, _ := strconv.Atoi(str)
	fmt.Println(num1)
	fmt.Println(strconv.Itoa(num))

	fmt.Println("________the second part of assignment________")
	fmt.Println()

	EuropeanVelocity := 120.4 * 3.6
	AmericanVelocity := 130 * 0.44704

	fmt.Println(fmt.Sprint(EuropeanVelocity) + " km/hour")
	fmt.Println(fmt.Sprint(AmericanVelocity) + " miles/hour")
}
