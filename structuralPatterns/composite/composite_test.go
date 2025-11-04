package composite

import "testing"

func TestComposite(t *testing.T) {
	root := Dir{
		Name: "structuralPatterns",
		Children: []Node{
			Dir{
				Name: "bridge",
				Children: []Node{
					File{Name: "bridge.go"},
					File{Name: "bridge_test.go"},
				},
			},
			Dir{
				Name: "composite",
				Children: []Node{
					File{Name: "composite.go"},
					File{Name: "composite_test.go"},
				},
			},
		},
	}

	root.Display("")
}
