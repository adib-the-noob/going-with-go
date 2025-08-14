package main

import "fmt"

type OrderStatus string 

const (
	Recieved OrderStatus = "recievedd"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing status to ", status)
}

func main() {
	changeOrderStatus(Delivered)
}