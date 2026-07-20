package main

import "fmt"

func ifElseExample() {
	// In Go we do not have ternary operators
	var a int = 22
	var b int = 3

	if a > 7 || b < 2 {
		fmt.Println("yes")
	} else if a == 3 && b == 4 {
		fmt.Println("No")
	} else {
		fmt.Println("Hi")
	}
}
