package memento

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Memento(t *testing.T) {
	manage := Manage{}
	text := TextNode{}
	text.SetState("1")
	manage.Save(text.Save())
	text.SetState("2")
	manage.Save(text.Save())
	state := text.GetState()
	fmt.Println(state)
	assert.True(t, manage.states[0].GetState() == "1")
	// 回退到第0步，状态是1
	assert.True(t, manage.Back(0).GetState() == "1")
}
