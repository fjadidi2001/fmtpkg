//package main
//
//import "fmt"
//
//func main() {
//	a := 353
//	b := "fjadidi"
//	const c = 45
//	d := 4 + 7i
//	var e float64 = 81.26
//
//	fmt.Printf("type a is :%T,it is :%+d\n", a, a)
//	fmt.Printf("type b is :%T,it is %s\n ", b, b)
//	fmt.Printf("type c is :%T,it is %d\n", c, c)
//	fmt.Printf("type is:%T,and it is %e\n", d, d)
//	fmt.Printf("type is %T and it is %g\n", e, e)
//}

// Golang program to illustrate the usage of
// fmt.Printf() function

// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	var str = "Geeksforgeeks"
	fmt.Printf("The string is %s \n", str)
	var num1 int = 21
	fmt.Printf("The decimal value is %d \n", num1)
	var num2 float32 = 7.786
	fmt.Printf("The floating point is %g \n", num2)
	var num3 int = 14
	fmt.Printf("The binary value of num3 is %b \n", num3)
	var num4 = 4 + 1i
	fmt.Printf("Scientific Notation of num4 : %e \n", num4)
}
