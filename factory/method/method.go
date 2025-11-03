package method

import "fmt"

type Pay interface {
	PayCode(price float64) string // 二维码地址
}

type AliPay struct{}

type WeiXinPay struct{}

type UnionPay struct{}

func (ali AliPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with AliPay\n", price)
	return "支付宝支付"
}

func (wei WeiXinPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with WeiXinPay\n", price)
	return "微信支付"
}

func (union UnionPay) PayCode(price float64) string {
	fmt.Printf("pay %#v with UnionPay\n", price)
	return "银联支付"
}

type PayFactory interface {
	CreatePay() Pay
}

// 创建两个厂
type AliPayFactory struct{}

func (AliPayFactory) CreatePay() Pay {
	// 专门处理支付宝支付的相关逻辑
	return AliPay{}
}

type WeiXinPayFactory struct{}

func (WeiXinPayFactory) CreatePay() Pay {
	// 专门处理微信支付的相关逻辑
	return WeiXinPay{}
}

// 新增银联支付
type UnionPayFactory struct{}

func (UnionPayFactory) CreatePay() Pay {
	return UnionPay{}
}
