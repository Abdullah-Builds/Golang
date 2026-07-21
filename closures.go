package main

import "fmt"

func increment() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {

	counter := increment()
	fmt.Println(counter())
	fmt.Println(counter())

}