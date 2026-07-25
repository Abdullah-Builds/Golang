package main

import "fmt"

//the use case of interface is that if we want payment via easypaisa and paypal
//  we have to make 2 instances in Payment struct 
//interfaces provides flexibility the methods in interface class all must be implemented by
//the struct and they are eligible to use the payment class

type paymenter interface {
	pay(amount float32)
	refund(amount float32, msg string)
}

type Payment struct {
	gateway paymenter
}

type PayPal struct{}

func (p PayPal) pay(amount float32) {
	fmt.Println("Paypal",amount)
}
func (p PayPal)refund(amount float32, msg string){
	fmt.Println(msg,amount)
}

type Easypaisa struct{}

func (p Easypaisa) pay(amount float32) {
	fmt.Println("Easypaisa",amount)
}
func (p Easypaisa)refund(amount float32, msg string){
	fmt.Println(msg,amount)
}

func main() {
    var EasypaisaObj = Easypaisa{}
	var PaymentObj = Payment{
		gateway: EasypaisaObj,
	}
	PaymentObj.gateway.pay(200)
	PaymentObj.gateway.refund(100,"Refunded $")
}