package bridge

import "fmt"

type Printer interface {
	PrintFile(file string) // 打印文件
}

type Epson struct{}

func (p *Epson) PrintFile(file string) {
	fmt.Printf("%s使用爱普生打印机打印文件", file)
}

type Hp struct{}

func (p *Hp) PrintFile(file string) {
	fmt.Printf("%s使用惠普打印机打印文件\n", file)
}

type Computer interface {
	Print(string)         // 打印
	SetPrinter(p Printer) // 设置打印机
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print(file string) {
	// 电脑调用打印机的打印方法
	m.printer.PrintFile("Mac电脑")
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (m *Windows) Print(file string) {
	// 电脑调用打印机的打印方法
	m.printer.PrintFile("Windows电脑")
}

func (m *Windows) SetPrinter(p Printer) {
	m.printer = p
}
