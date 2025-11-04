package facade

import "testing"

func TestFacade(t *testing.T) {
	inventory := Inventory{}
	inventory.Deduction()

	pay := Pay{}
	pay.Pay()

	logistics := Logistics{}
	logistics.Shipping()
}

func TestFacade_Order(t *testing.T) {
	o := NewOrder()
	o.PlaceOrder()
}
