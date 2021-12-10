package main

import "fmt"

func main() {

	week := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	fmt.Println(week)

	workDays := week[1:6]
	fmt.Println(workDays)

	// weekEnd := append(week[0:1])
	// weekEnd2 := append(week[6:7])
}
