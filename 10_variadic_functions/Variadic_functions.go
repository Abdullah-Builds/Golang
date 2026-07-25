package main

import "fmt"

func AcceptIntAsmany(num ...int) {
	fmt.Println(num)
}

func AcceptAny(val ... interface{}){
	fmt.Println(val...)
}

func main() {
    
	AcceptIntAsmany(1,2,3,4)
	AcceptAny("khan",1,true,10.22)

	var num = []int {1,2,3,4}
    AcceptIntAsmany(num...)
}