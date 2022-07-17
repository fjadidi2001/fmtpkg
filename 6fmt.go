package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 = 23
	const num2 = 10
	var numberi = "Hello means salam"
	var work = " Nice job"
	var sima = 'r'
	fmt.Scanln(&numberi) //pointer format
	num3, err := strconv.Atoi(numberi)
	if err != nil {
		fmt.Println("ye number vared kon")
		return
	}
	num4 := -4 + 3i
	num5 := 25.36

	if num1 < num2 {
		fmt.Printf("this is %c", num2)
	} else {
		fmt.Printf("base 8 num 1 is %o\n", num1)
		fmt.Printf("base 8 num 1 with prefix is %O\n", num1)
		fmt.Printf("base 16 num 1 is %x\n", num1)
		fmt.Printf("base 16 num2 is %X\n\n", num2)
		fmt.Printf("base 16 num3 is %X\n", num3)
		fmt.Printf("string is %s\n", work)
		fmt.Printf("string format with quoted %q\n", work)
		fmt.Printf("format for Character %c\n", sima)
		fmt.Printf("format for unicode: %U\n", num2)
		fmt.Printf("format for scientific notation%E\n ", num4)
		fmt.Printf("float format %.5f\n", num5)
	}
}
