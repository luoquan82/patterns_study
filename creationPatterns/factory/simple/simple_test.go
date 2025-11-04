package simple

import (
	"testing"
)

func TestNewPayCode(t *testing.T) {
	pay := NewPayCode(AliPayType)
	pay.PayCode(66.2)

	pay = NewPayCode(WeiXinPayType)
	pay.PayCode(66.7)
}
