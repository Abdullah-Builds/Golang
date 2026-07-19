package main

import "fmt"

const name = "khan"
//const name := "khan"  //error 
func main() {

	const age = 30
	const Age int = 40

	fmt.Println(Age,age)

	//multiple constants
	const (
		port = 400
		uni = "fast"
	)

	fmt.Println(port,uni)
}