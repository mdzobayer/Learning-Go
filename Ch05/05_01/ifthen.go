package main

import "fmt"

func main() {

	//var x float64 = -5
	var result string

	if x := 42; x < 0 {
		result = "Less than zero"
	} else {
		result = "Greater than or equal to zero"
	}
	fmt.Println("Result is:", result)
	//fmt.Println("X is:", x)	// x is undefined

}
