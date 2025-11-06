package visitor

import (
	"fmt"
	"testing"
)

func Test_Visitor(t *testing.T) {
	document := &Document{}
	document.AddElement(&Text{content: "你好"})
	document.AddElement(&Text{content: "XXX"})
	document.AddElement(&Image{src: "abc.png"})

	pdf := &PDFVisitor{}
	document.Accept(pdf)

	fmt.Println("=======================================")

	html := &HTMLVisitor{}
	document.Accept(html)
}
