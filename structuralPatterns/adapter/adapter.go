package adapter

import "fmt"

type AliPay struct{}

func (a AliPay) GetPayPage(price float64) string {
	return fmt.Sprintf("支付宝支付链接: XXXXXX,支付金额%f", price)
}

type AliPayAdapter struct {
	aliPay *AliPay
}

func (a AliPayAdapter) PayPage(price float64) string {
	return a.aliPay.GetPayPage(price)
}

type WeiXinPay struct{}

func (w WeiXinPay) PayPage(price float64) string {
	return fmt.Sprintf("微信支付链接: XXXXXX,支付金额%f", price)
}

type PayI interface {
	PayPage(price float64) string
}

func PayPage(pi PayI, price float64) string {
	return pi.PayPage(price)
}
