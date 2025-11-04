package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	req := Req{}
	reqDecorator := LogReqDecorator{req: req}
	reqDecorator.Handler("http://www.baidu.com")

	monitorDecorator := MonitorDecorator{req: reqDecorator}
	monitorDecorator.Handler("http://www.baidu.com")
}
