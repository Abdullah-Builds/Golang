package main

import "fmt"

func main() {

	var m = map[string]int{}
	m["val1"] = 8
	m["val2"] = 22
	fmt.Println(m["val1"])

	mapping := make(map[string]string)
	mapping["value"] = "khan"
	fmt.Println(mapping)

	delete(m, "val1")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
}
