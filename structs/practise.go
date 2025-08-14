package main

import (
	"fmt"
	"time"
)

type payment struct {
	total_amount float32
	quantity int
}

type order struct {
	id string
	amount float32
	status string
	createAt time.Time
	payment
}

// reciever type
func (o *order) changeDelStatus(status string) {
	o.status = status
}

func createOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id: id,
		amount: amount,
		status: status,		
	}
	return &myOrder
}

func main() {
	newOrder := order{
		id: "12",
		amount: 12,
		createAt: time.Now(),
		payment: payment{
			total_amount: 12.12,
			quantity: 12,
		},
	}
	fmt.Println(newOrder.payment)
}

