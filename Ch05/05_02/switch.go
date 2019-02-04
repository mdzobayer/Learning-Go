package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""
	dow = 7

	switch dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	case 3:
		result = "It's Tuesday"
	case 4:
		result = "It's Wednessday"
	case 5:
		result = "It's Thursday"
	case 6:
		result = "It's Friday"
	default:
		result = "It's Satarday"

	}

	fmt.Println("Day :", dow, "Today is :", result)

	x := 5
	switch {
	case x < 0:
		result = "Less then zero"
	case x == 0:
		result = "Equal to zero"
	default:
		result = "Grater then zero"
	}
	fmt.Println(result)

}
