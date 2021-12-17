package main

import (
	"fmt"
)

func main() {
	issuedEdition := map[string]map[string][]string{
		"aSavchenko": {
			"Java.Professional library": {"Volume 1. Basics.", "Volume 2. Advanced Software"},
		},
		"rStokolyas": {
			"Cloud Go": {"first edition"},
		},
		"dShvets": {
			"Ruby": {"Ruby for Romantics. From beginner to professional", "Ruby for Romantics - Second Edition"},
		},
	}

	numberOfReaders := len(issuedEdition)

	fmt.Printf("number of readers : %v\n", numberOfReaders)

	for k, v := range issuedEdition {
		fmt.Print(k, ": ")
		for subK, subV := range v {
			fmt.Print(subK, ": ", len(subV), " edition")
		}
		fmt.Println()
	}

}
