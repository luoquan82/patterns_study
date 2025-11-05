package templateMethod

import "testing"

func Test_TemplateMethod(t *testing.T) {
	coffee := Coffee{}
	tea := Tea{}

	MakeTemplate(&coffee)
	MakeTemplate(&tea)
}
