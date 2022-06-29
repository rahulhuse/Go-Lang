package main

import "fmt"

func main() {

	ans := 1

	switch {
	case ans > 0:
		fmt.Println("Number is greater than zero")
	case ans < 0:
		fmt.Println("Number is less than zero")
	default:
		fmt.Println("Number is zero")

	}
}
