package abstract

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	aliPayFactory := AliPayFactory{}
	aliPay := aliPayFactory.CreatePay()
	aliPay.PayCode(11)
	err := aliPayFactory.CreateRefund().Refund("11")
	if err != nil {
		t.Error(err)
	}
}
