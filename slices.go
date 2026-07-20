package main

import "fmt"

func slicesExample() {
	// unintialized slice is nil
	var num1 []int
	fmt.Println(num1 == nil)

	var num = make([]int, 0, 3) // ([]int , initial length, capacity )
	num = append(num, 23)
	fmt.Println(num)
	fmt.Println(cap(num), len(num))

	num2 := []int{}
	num2 = append(num2, 2)
	fmt.Println(num2)
	fmt.Println(cap(num2), len(num2))

	var nums3 = []int{1, 2, 3, 4}
	fmt.Println(nums3[0:2])
}
