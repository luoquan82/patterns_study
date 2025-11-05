package interpreter

import (
	"fmt"
	"testing"
)

func Test_Interpreter(t *testing.T) {
	tmpl := `Hello, {{ Name }}! You are {{ Age}} years old.`
	template := ParseTemplate(tmpl)
	fmt.Println(template)
	res := template.Interpreter(&Context{
		Data: map[string]any{
			"Name": "zhangsan",
			"Age":  18,
		},
	})
	fmt.Println(res)
}
