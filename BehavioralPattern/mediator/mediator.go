package mediator

import "fmt"

type Obj interface {
	SendMsg(string)
	RecMsg(string)
}

type Mediator interface {
	SendMsg(msg string, to Obj)
}

type User struct {
	Name     string
	mediator Mediator
}

func (u User) SendMsg(msg string) {
	fmt.Printf("用户 %s 发了消息 %s\n", u.Name, msg)
	u.mediator.SendMsg(msg, u)
}

func (u User) RecMsg(msg string) {
	fmt.Printf("用户 %s 接收了消息 %s\n", u.Name, msg)

}

type ChatRoom struct {
	Users []User
}

func (c *ChatRoom) SendMsg(msg string, user Obj) {
	for _, u := range c.Users {
		if u != user {
			u.RecMsg(msg)
		}
	}
}

func (c *ChatRoom) Register(user User) {
	c.Users = append(c.Users, user)
}
