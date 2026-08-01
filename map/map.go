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

	//if key is not present go does not gives error it gives the value type by default
	// int -> 0
	// bool ->false
	//string -> " "/ empty string

// Safe key Validation Technique 
	_, status := m["val1"]

	if status {
		fmt.Println("Exists")
	}else{
		fmt.Println("Not Exists")
	}
}
