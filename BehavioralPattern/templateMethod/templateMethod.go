package templateMethod

import "fmt"

type Template interface {
	BoilWater()        // 烧水
	Brew()             // 冲泡
	AddSugar()         // 加糖
	HasAddSugar() bool // 是否加糖
}

type Coffee struct{}

func (Coffee) BoilWater() {
	fmt.Println("咖啡烧水")
}

func (Coffee) Brew() {
	fmt.Println("冲泡咖啡")
}

func (Coffee) AddSugar() {
	fmt.Println("咖啡加糖")
}

func (Coffee) HasAddSugar() bool {
	return true
}

type Tea struct{}

func (Tea) BoilWater() {
	fmt.Println("烧水煮茶")
}

func (Tea) Brew() {
	fmt.Println("冲泡茶")
}

func (Tea) AddSugar() {
	fmt.Println("加糖")
}

func (Tea) HasAddSugar() bool {
	return false
}

func MakeTemplate(tmp Template) {
	tmp.BoilWater()
	tmp.Brew()
	if tmp.HasAddSugar() {
		tmp.AddSugar()
	}
}
