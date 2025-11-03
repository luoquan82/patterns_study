package simple

import "fmt"

type Pay interface {
	PayCode(price float64) string // 二维码地址
}

type AliPay struct{}

type WeiXinPay struct{}

func (ali AliPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with AliPay\n", price)
	return "支付宝支付"
}

func (wei WeiXinPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with WeiXinPay\n", price)
	return "微信支付"
}

type PayType int8

const (
	AliPayType PayType = iota + 1
	WeiXinPayType
)

func NewPayCode(t PayType) Pay {
	switch t {
	case AliPayType:
		return AliPay{}
	case WeiXinPayType:
		return WeiXinPay{}
	}
	return nil
}
