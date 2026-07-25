package main

import "fmt"

type OrderStatus int
type OrderMsg string

const (
	Received  OrderMsg = "Received"
	Cancelled OrderMsg = "Cancelled"
	Delayed   OrderMsg = "Delayed"
)

const (
	received OrderStatus = iota
	cancelled
	delayed
)

func CheckOrderStatus(status OrderStatus) {
	fmt.Println(status)
}
func OrderMsgStatus(status OrderMsg) {
	fmt.Println(status)
}
func main() {
	CheckOrderStatus(received)
	OrderMsgStatus(Delayed)
}
