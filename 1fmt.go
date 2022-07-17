package main

import "fmt"

func main() {
	a, b := 12, 23

	if a > b {
		fmt.Printf("binary %b\\%b", a, b)
		return

	}
	fmt.Printf("binary: %b\\%b\n", a, b)

}
