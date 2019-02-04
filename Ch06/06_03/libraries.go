package main

import (
	"fmt"
)

func main() {
	n1, l1 := fullName("Zobayer", "Mahmud")

	fmt.Printf("Length : %v Name : %v\n", l1, n1)

	n1, l1 = fullName("Zobayer", "Mahmud")

	fmt.Printf("New Length : %v Name : %v\n", l1, n1)
}

func fullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)

	return full, length
}

func fullNameNackedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)

	return
}
