package main

import "fmt"

func add(a *int) {
	*a = 20
	fmt.Println(*a)
}
func main() {
	var a int = 30
	add(&a)
	fmt.Println(a)

}