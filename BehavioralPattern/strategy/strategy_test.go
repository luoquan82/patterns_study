package strategy

import "testing"

func Test_Strategy(t *testing.T) {
	aliPay := &AliPay{}
	weiPay := &WeiPay{}

	strategy := &PayStrategy{}
	strategy.SetPay(aliPay)
	strategy.Pay(12)
	strategy.SetPay(weiPay)
	strategy.Pay(13)
}
