package main

import "fmt"

func Print[T int | string](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

func PrintAny[T any](arr[]T){
	for _, item := range arr {
		fmt.Println(item)
	}
}
func main() {
    var arr1 = []string{"js", "ts"}
	var arr2 = []bool{true, false}
	
	Print(arr1)
	PrintAny(arr2)
}