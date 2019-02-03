package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("P is nil")
	}

	var v int
	v = 42
	p = &v

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("P is nil")
	}

	var value1 float64 = 42.13

	pointer := &value1

	*pointer = *pointer * 5

	if p != nil {
		fmt.Println("Value of p:", *pointer)
	} else {
		fmt.Println("Pointer is nil")
	}
}
