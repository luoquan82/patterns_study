package abstract

import "fmt"

type Pay interface {
	PayCode(price float64) string // 二维码地址
}

type AliPay struct{}

type WeiXinPay struct{}

// 退款
type Refund interface {
	Refund(no string) error
}

type AliRefund struct{}

func (AliRefund) Refund(no string) error {
	fmt.Printf("订单号 %#v,支付宝退款\n", no)
	return nil
}

type WeiXinRefund struct{}

func (WeiXinRefund) Refund(no string) error {
	fmt.Printf("订单号 %#v,微信退款\n", no)
	return nil
}

func (ali AliPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with AliPay\n", price)
	return "支付宝支付"
}

func (wei WeiXinPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with WeiXinPay\n", price)
	return "微信支付"
}

type PayFactory interface {
	CreatePay() Pay
	CreateRefund() Refund
}

// 创建两个厂
type AliPayFactory struct{}

func (AliPayFactory) CreatePay() Pay {
	// 专门处理支付宝支付的相关逻辑
	return AliPay{}
}

func (AliPayFactory) CreateRefund() Refund {
	return AliRefund{}
}

func (WeiXinRefund) CreateRefund() Refund {
	return WeiXinRefund{}
}

type WeiXinPayFactory struct{}

func (WeiXinPayFactory) CreatePay() Pay {
	// 专门处理微信支付的相关逻辑
	return WeiXinPay{}
}
