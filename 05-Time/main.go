package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Module")

	presentTime := time.Now()
	fmt.Println("Present Time :", presentTime)

	// For formating the date or time we need to give in "01-02-2006 Monday 15:04:05" format.
	// These should be same
	// Launch Date for GOLANG
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.August, 15, 12, 12, 2, 2, time.UTC)
	fmt.Println("createdDate :", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}
