package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	weiXinPay := PayPage(WeiXinPay{}, 16.7)
	fmt.Println(weiXinPay)

	aliPay := PayPage(AliPayAdapter{aliPay: &AliPay{}}, 16.8)
	fmt.Println(aliPay)
}
