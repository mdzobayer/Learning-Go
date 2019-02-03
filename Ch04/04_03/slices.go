package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}

	fmt.Println("Size of colors:", len(colors))
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[2:len(colors)])  // Remove first item
	colors = append(colors[:len(colors)-1]) // Remove last item
	fmt.Println(colors)

	// use function make()
	numbers := make([]int, 10, 10)
	fmt.Println(numbers)

	numbers[0] = 285
	numbers[1] = 47
	numbers[2] = 96
	numbers[3] = 85
	numbers[4] = 21
	fmt.Println(cap(numbers))
	fmt.Println(numbers)

	numbers = append(numbers, 49)
	numbers = append(numbers, 89)
	fmt.Println(cap(numbers))
	numbers = append(numbers, 49)
	//numbers = append(numbers, 89)
	//numbers = append(numbers, 49)
	//numbers = append(numbers, 89)
	fmt.Println(cap(numbers))
	sort.Ints(numbers)
	fmt.Println(numbers)

}
