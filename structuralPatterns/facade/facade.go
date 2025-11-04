package facade

import "fmt"

type Inventory struct{}

func (k Inventory) Deduction() {
	fmt.Println("扣库存")
}

type Pay struct{}

func (p Pay) Pay() {
	fmt.Println("支付")
}

// Logistics 物流
type Logistics struct{}

func (l Logistics) Shipping() {
	fmt.Println("发货")
}

type Order struct {
	inventory *Inventory
	pay       *Pay
	logistics *Logistics
}

func NewOrder() *Order {
	return &Order{
		inventory: &Inventory{},
		pay:       &Pay{},
		logistics: &Logistics{},
	}
}

func (o Order) PlaceOrder() {
	o.inventory.Deduction()
	o.pay.Pay()
	o.logistics.Shipping()
}
