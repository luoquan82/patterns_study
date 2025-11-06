package visitor

import "fmt"

type Element interface {
	Accept(v Visitor)
}
type Text struct {
	content string
}

func (t *Text) Accept(v Visitor) {
	v.VisitText(t)
}

type Image struct {
	src string
}

func (img *Image) Accept(v Visitor) {
	v.VisitImage(img)
}

type Visitor interface {
	VisitText(text *Text)
	VisitImage(image *Image)
}

type PDFVisitor struct{}

func (p PDFVisitor) VisitText(text *Text) {
	fmt.Printf("PDF 获取文本内容: %s\n", text.content)
}

func (p PDFVisitor) VisitImage(image *Image) {
	fmt.Printf("PDF 获取图片的src: %s\n", image.src)
}

type HTMLVisitor struct{}

func (p HTMLVisitor) VisitText(text *Text) {
	fmt.Printf("HTML 获取文本内容: %s\n", text.content)
}

func (p HTMLVisitor) VisitImage(image *Image) {
	fmt.Printf("HTML 获取图片的src: %s\n", image.src)
}

type Document struct {
	elements []Element
}

func (d *Document) AddElement(element Element) {
	d.elements = append(d.elements, element)
}

func (d *Document) Accept(visitor Visitor) {
	for _, element := range d.elements {
		element.Accept(visitor)
	}
}
