package method

import "testing"

func TestMethod(t *testing.T) {
	aliPayFactory := AliPayFactory{}
	aliPay := aliPayFactory.CreatePay()
	aliPay.PayCode(16.8)

	weiXinFactory := WeiXinPayFactory{}
	weiXinPay := weiXinFactory.CreatePay()
	weiXinPay.PayCode(16.9)

	unionFactory := UnionPayFactory{}
	unionPay := unionFactory.CreatePay()
	unionPay.PayCode(17.0)
}
