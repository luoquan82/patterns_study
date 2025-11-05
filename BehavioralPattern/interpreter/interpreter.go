package interpreter

import (
	"fmt"
	"strings"
)

type Context struct {
	Data map[string]any
}

type Node interface {
	Interpreter(*Context) string
}

type TextNode struct {
	Content string
}

func (t *TextNode) Interpreter(ctx *Context) string {
	return t.Content
}

type Template struct {
	tree []Node
}

type VarNode struct {
	Key string
}

func (t *VarNode) Interpreter(ctx *Context) string {
	v, ok := ctx.Data[t.Key]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v)
}

func ParseTemplate(tmpl string) *Template {
	var template = new(Template)
	var index = 0
	for {
		startIndex := strings.Index(tmpl[index:], "{{")
		if startIndex == -1 {
			template.tree = append(template.tree, &TextNode{
				Content: tmpl[index:],
			})
			break
		}
		template.tree = append(template.tree, &TextNode{
			Content: tmpl[index : index+startIndex],
		})

		endIndex := strings.Index(tmpl[index+startIndex:], "}}")
		if endIndex == -1 {
			break
		}

		key := tmpl[index+startIndex+2 : index+startIndex+endIndex]
		key = strings.TrimSpace(key)
		template.tree = append(template.tree, &VarNode{
			Key: key,
		})
		index += startIndex + endIndex + 2
	}
	return template
}

func (t *Template) Interpreter(ctx *Context) string {
	var result string
	for _, node := range t.tree {
		result += node.Interpreter(ctx)
	}
	return result
}
