package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func returnMany() (string, string, int) {
	return "js", "go", 1
}
func main() {
	var res = sum(2, 3)
	fmt.Println(res)

	fmt.Println(returnMany())
}
