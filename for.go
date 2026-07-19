package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {

		if i == 2 {
			continue
		}
		
		fmt.Println(i)
	}


	j := 0
	for j <=3 {

		if j==2 {
			break
		}
		j++
	}

	
	for k :=range 10 {
      fmt.Println(k)
	}
}