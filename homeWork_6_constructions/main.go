package main

import (
	"fmt"
)

func contains(s []string, x string) bool {
	for _, v := range s {
		if x == v {
			return true
		}
	}
	return false
}

func getMax(p ...int) int {
	var res int
	for _, v := range p {
		res += v
	}
	return res
}

func main() {
	s := []string{"a", "b", "c"}
	x := "b"
	fmt.Println("contains = ", contains(s, x))
	fmt.Println("getMax = ", getMax(3, 6, 5))
}
