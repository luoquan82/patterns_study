package composite

import "fmt"

type Node interface {
	Display(indent string)
}

type File struct {
	Name string
}

type Dir struct {
	Name     string
	Children []Node
}

func (f File) Display(indent string) {
	fmt.Println(indent + f.Name)
}

func (d Dir) Display(indent string) {
	fmt.Println(indent + d.Name)
	for _, child := range d.Children {
		child.Display(indent + "  ")
	}
}
