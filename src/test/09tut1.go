package main

import "fmt"

func main() {

	age := 10

	if age >= 19 {

		fmt.Println("You can ride")
	} else {
		fmt.Println("You can not ride")
		fmt.Println("Wait for ", 18-age, "years")
	}

}
