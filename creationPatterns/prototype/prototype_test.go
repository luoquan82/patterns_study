package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	s1 := Student{Name: "luoquan", Age: 40}

	s2 := s1.Clone().(*Student)
	s2.Name = "zhangsan"
	fmt.Println(s1)
	fmt.Println(s2)
}
