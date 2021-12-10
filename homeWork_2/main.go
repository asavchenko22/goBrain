package main

import (
	"fmt"
	"math"
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

	type EuropeanVelocity float64
	type AmericanVelocity float64

	var ev EuropeanVelocity = 120.4 * 3.6
	var av AmericanVelocity = 130 * 2.2369362920544

	fmt.Println(ev, " km/hour")
	fmt.Println(math.Round(float64(av*100))/100, " miles/hour")
}
