package main

import (
	"fmt"
)

type Personal_Details struct {
	name    string
	age     int
	citizen bool
}
//There are no default constructors so we can make it like this 
func Constructor(name string, age int, citizen bool)Personal_Details{
	var MyObj = Personal_Details{
		name : name ,
		age : age,
		citizen: citizen,
	}
	return MyObj
}
// these funcs are mapped to struct just like behaves as a class in C++
func(P *Personal_Details) ModifyObj(citizen bool){
     P.citizen = citizen
}
func(P Personal_Details) getAge() int {
	return P.age
}

func main() {
//=>1
	var Obj2 = Constructor("izhan",122,true)
	Obj2.name = "abd"
	fmt.Println(Obj2)

//=>2
	var Obj = Personal_Details{
		name:    "Khan",
		age:     22,
		citizen: true,
	}

	fmt.Println(Obj)
	Obj.ModifyObj(false)
	fmt.Println(Obj.getAge())

//=>3
     var language = struct{
		lang string;
		native bool
	 }{"Urdu",false}
     fmt.Println(language)
}