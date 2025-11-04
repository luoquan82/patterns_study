package command

import "testing"

func Test_Command(t *testing.T) {
	queue := NewTaskQueue()
	queue.AddCommand(&PrintCommand{Content: "你好"})
	queue.AddCommand(&SendEmail{Content: "你好", To: "XXX.qq.com"})
	queue.Command()
}
