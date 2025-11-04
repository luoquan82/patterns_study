package bridge

import "testing"

func TestBridge(t *testing.T) {
	w := Windows{}
	hp := Hp{}
	w.SetPrinter(&hp)
	w.Print("")

	m := Mac{}
	e := Epson{}
	m.SetPrinter(&e)
	m.Print("")
}
