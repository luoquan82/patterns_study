package mediator

import "testing"

func Test_Mediator(t *testing.T) {
	room := ChatRoom{}
	u1 := User{Name: "zhangsan", mediator: &room}
	u2 := User{Name: "lisi", mediator: &room}
	u3 := User{Name: "wangwu", mediator: &room}

	room.Register(u1)
	room.Register(u2)
	room.Register(u3)

	u1.SendMsg("你好啊")
	u2.SendMsg("吃了吗?")
	u3.SendMsg("我吃了")
}
