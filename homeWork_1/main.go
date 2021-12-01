package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("Hello mr.X! today is 02.01.2006 15:04"))
}
