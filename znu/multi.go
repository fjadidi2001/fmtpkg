package znu

import "fmt"

func Multi(num1, num2 float64) float64 {
	var mul = num1 * num2
	fmt.Printf("type of mul is %T,the decimal value is %g\n", mul, mul)
	return mul
}
