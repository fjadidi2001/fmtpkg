package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 = 23
	const num2 = 10
	var myString = "Hello means salam"
	var work = " Nice job"
	fmt.Scanln(&myString) //be sorat pointer ersal mikonim
	num3, err := strconv.Atoi(myString)
	if err != nil {
		fmt.Println("ye number vared kon")
		return
	}
	num4 := -4 + 3i
	if num1 < num2 {
		fmt.Printf("this is %c", num2)
	} else {
		fmt.Printf("base 8 num 1 is %o\n", num1)
		fmt.Printf("base 8 num 1 with prefix is %O\n", num1)
		fmt.Printf("base 16 num 1 is %x\n", num1)
		fmt.Printf("base 16 num2 is %X\n\n", num2)
		fmt.Printf("base 16 num3 is %X\n", num3)
		fmt.Printf("string is %s\n", myString)
		fmt.Printf("try to use %q\n", work)
		fmt.Printf("format for Character %c\n", work)
		fmt.Printf("format for unicode: %U\n", num2)
		fmt.Printf("format for scientific notation%E\n ", num4)
		fmt.Printf("float")
	}
}
