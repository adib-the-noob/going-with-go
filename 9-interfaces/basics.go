package main
import (
	"fmt"
)

type payment struct {
	gateway stripe  	
}

type razorPay struct {

}

type stripe struct {

}

func (r razorPay) pay (amount float32) {
	// make payment
	fmt.Println("Making payment with Razor Pay..", amount)
}

func (s stripe) pay (amount float32) {
	// make payment
	fmt.Println("Making payment with STRIPE Pay..", amount)
}


func (p payment) makePayment(amount float32) {
	p.gateway.pay(12);
}

func main() {
	newPayment := payment{}
	newPayment.makePayment(120)
}
