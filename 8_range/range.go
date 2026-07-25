package main

import "fmt"

func main() {

	var slice = []int{1, 2, 3, 4}

	for index, val := range slice {
		fmt.Println(index,val)
	}

	var mapp = map[string]int{"one":1, "two":2}
    for key, val := range mapp{
		fmt.Println(key,val)
	}
}