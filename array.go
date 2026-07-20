package main

import "fmt"

func main() {

	// Use arrays when we have fixed size

	var num = [3]int{1, 2, 3}
	Num := [3]int{1, 2, 3}
	fmt.Println(num[0], Num)

	var name [3]string
	name[0] = "Go"
	fmt.Println(len(name))

	age := [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(age)
}
