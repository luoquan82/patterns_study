package builder

import (
	"fmt"
	"testing"
)

func TestBao(t *testing.T) {
	bao := NewBao()
	boss := NewBoss(bao)
	fmt.Println(boss.GetHouse())
}
