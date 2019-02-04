package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	numbers := []int{5, 99, 4, 3, 48, 74}
	fmt.Println(colors)

	sum = 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum is :", sum)

	for x := range numbers {
		fmt.Println("Number: ", numbers[x])
	}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for color := range colors {
		fmt.Println(colors[color])
	}

	// 	sum = 1
	// 	for sum < 1000 {
	// 		sum += sum
	// 		if sum == 16 {
	// 			continue
	// 		}
	// 		fmt.Println(sum)
	// 		if sum > 200 {
	// 			goto endofprogram
	// 		}
	// 		if sum > 500 {
	// 			break
	// 		}
	// 	}

	// endofprogram:
	// 	fmt.Println("End of Program")

}
