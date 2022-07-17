package znu

import "fmt"

func Plus(num1, num2 float64) float64 {
	var sum = num1 + num2
	fmt.Printf("sumation is :%g,type of sum is %T\n", sum, sum)
	return sum
}
