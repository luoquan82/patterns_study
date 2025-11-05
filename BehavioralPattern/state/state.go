package state

import "fmt"

type State interface {
	Switch(ctx *Context)
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Switch() {
	c.state.Switch(c)
}

type OnState struct{}

func (OnState) Switch(ctx *Context) {
	fmt.Println("开关关闭")
	ctx.SetState(&OffState{})
}

type OffState struct{}

func (OffState) Switch(ctx *Context) {
	fmt.Println("开关打开")
	ctx.SetState(&OnState{})
}
