package fundamental

import "fmt"

// Struct Payment disebut sbg objek
type Payment struct {
	ID       int
	PoNumber string
	Type     string
}

func Struct() {

	println("In Struct File")

	var user Payment

	user.ID = 1
	user.PoNumber = "ID_001"
	user.Type = "VA"

	buyer := Payment{
		ID:       2,
		PoNumber: "ID_002",
		Type:     "E-Wallet",
	}

	fmt.Println("user:", user)
	fmt.Println("buyer:", buyer)

	msg := buyer.sendPayment()
	println(msg)

	println()
}

// Struct Method = Function yg menempel/dimiliki objek/struct
func (p *Payment) sendPayment() string {
	return "Processing Payment"
}
