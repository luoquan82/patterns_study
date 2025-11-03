package basic

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	d1 := GetInstance()
	d2 := GetInstance()
	d3 := GetInstance()

	if d1 != d2 {
		t.Errorf("Singleton different.")
	}

	if d3 != d2 {
		t.Errorf("Singleton different.")
	}

	fmt.Printf("%#p %#p %#p\n", d1, d2, d3)
}
