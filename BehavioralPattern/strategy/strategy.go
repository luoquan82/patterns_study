package strategy

import "fmt"

type Pay interface {
	Pay(money float64)
}

type AliPay struct{}

func (AliPay) Pay(money float64) {
	fmt.Printf("使用支付宝支付 %f\n", money)
}

type WeiPay struct{}

func (WeiPay) Pay(money float64) {
	fmt.Printf("使用微信支付 %f\n", money)
}

type PayStrategy struct {
	pay Pay
}

func (p *PayStrategy) SetPay(pay Pay) {
	p.pay = pay
}

func (p *PayStrategy) Pay(money float64) {
	p.pay.Pay(money)
}
