package main

import "fmt"

func main() {
	doSomething()
	//var sum int
	sum := addValue(10, 5)

	fmt.Println("Sum is :", sum)

	sum = addAllValue(5, 8, 6, 4, 7)
	fmt.Println("New sum is :", sum)
}
func doSomething() {
	fmt.Println("Doing Something")
}

//func addValue(value1 int, value2 int) int { // work fine
func addValue(value1, value2 int) int {
	return value1 + value2
}

func addAllValue(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}

	//fmt.Printf("%T\n", values)

	return sum
}
