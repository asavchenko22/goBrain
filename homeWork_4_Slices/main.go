package main

import "fmt"

func main() {

	week := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	fmt.Printf("week: %v\n", week)

	workDays := make([]string, 5)
	copy(workDays, week[1:6])
	fmt.Printf("workDays: %v\n", workDays)

	week = append(week[:1], week[6:]...)
	fmt.Printf("week: %v\n", week)

	combinedWeek := append(week[0:1:1], workDays...)
	combinedWeek = append(combinedWeek, week[1:]...)

	fmt.Printf("combinedWeek: %v\n", combinedWeek)
	fmt.Printf("week: %v\n ", week)
	fmt.Printf("workDays: %v\n ", workDays)
}
