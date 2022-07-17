package main

import (
	"fmt"
	"fmtpkg/znu"
)

func main() {
	a := znu.Multi(25.26, 87.29)
	b := znu.Plus(78.36, 45.23)
	fmt.Println(a, b)
}