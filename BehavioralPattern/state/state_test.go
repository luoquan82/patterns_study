package state

import "testing"

func Test_State(t *testing.T) {
	c := &Context{
		state: &OffState{},
	}
	c.Switch()
	c.Switch()
}
