package decorator

import (
	"fmt"
	"time"
)

type ReqI interface {
	Handler(url string) string
}

type Req struct{}

func (m Req) Handler(url string) string {
	fmt.Println("请求 " + url)
	return ""
}

// LogReqDecorator 日志装饰器
type LogReqDecorator struct {
	req ReqI
}

func (l LogReqDecorator) Handler(url string) string {
	fmt.Println("日志记录前")
	res := l.req.Handler(url)
	fmt.Println("日志记录后")
	return res
}

type MonitorDecorator struct {
	req ReqI
}

func (m MonitorDecorator) Handler(url string) string {
	s := time.Now()
	res := m.req.Handler(url)
	sub := time.Since(s)
	fmt.Printf("耗时%s", sub)
	return res
}
