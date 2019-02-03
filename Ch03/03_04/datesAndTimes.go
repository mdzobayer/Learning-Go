package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("The month is ", t.Month())
	fmt.Println("The Day is ", t.Day())
	fmt.Println("The weekday is ", t.Weekday())
	fmt.Println("The hour is ", t.Hour())

	fmt.Println("The month is ", now.Month())
	fmt.Println("The Day is ", now.Day())
	fmt.Println("The weekday is ", now.Weekday())
	fmt.Println("The hour is ", now.Hour())

	tomorrow := now.AddDate(0, 0, 1)

	fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(), tomorrow.Month(),
		tomorrow.Day(), tomorrow.Year())

	longFormate := "Thursday, January, 31, 2019"

	fmt.Println("Tomorrow is", tomorrow.Format(longFormate))
	shortFormate := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormate))
}
