package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var s string
	//fmt.Scanln(&s)
	// fmt.Println("You typed:", s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter line: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("Line is: ", str)

	// fmt.Print("Enter text: ")
	// str, _ = reader.ReadString(' ')
	// fmt.Println("Text is: ", str)

	fmt.Print("Enter Number: ")
	str, _ = reader.ReadString('\n')
	// f, err := strconv.ParseFloat(str, 64)
	//f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	//g, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err == nil {
		fmt.Println("f is :", f)
		//fmt.Println("f is :", f, "g is :", g)
	} else {
		fmt.Println(err)
	}
}
