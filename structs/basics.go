package main

import (
	"fmt"
	"time"
)

type order struct {
	id string
	amount float32
	status string
	createAt time.Time

}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return  o.amount
}

func main() {
	order := order{
		id: "12",
		amount: 12.33,
		status: "yes",
	}
	order.createAt = time.Now()
	fmt.Println(order)
	order.changeStatus("delivered")
	fmt.Println(order)
	fmt.Println(order.getAmount())
}
