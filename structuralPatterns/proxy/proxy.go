package proxy

import "fmt"

type Log struct{}

func (Log) Info(content string) {
	fmt.Printf("%#v 日志落库\n", content)
}

type ProxyLog struct {
	log *Log
}

func (p *ProxyLog) Info(content string) {
	// 访问前
	// 延迟初始化
	if p.log == nil {
		p.log = &Log{}
	}
	p.log.Info(content)
	// 访问后
	fmt.Println("记录 调用了 info")
}
