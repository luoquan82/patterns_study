package command

import "fmt"

type Command interface {
	Execute()
}

type PrintCommand struct {
	Content string
}

func (p PrintCommand) Execute() {
	fmt.Printf("打印消息: %s\n", p.Content)
}

type SendEmail struct {
	To      string
	Content string
}

func (s SendEmail) Execute() {
	fmt.Printf("发送邮件给:%s,内容:%s\n", s.To, s.Content)
}

type TaskQueue struct {
	Queue []Command
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{}
}

func (t *TaskQueue) AddCommand(command Command) {
	t.Queue = append(t.Queue, command)
}

func (t *TaskQueue) Command() {
	for _, command := range t.Queue {
		command.Execute()
	}
}
