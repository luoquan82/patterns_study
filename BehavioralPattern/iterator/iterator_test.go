package iterator

import "testing"

func Test_Iterator(t *testing.T) {
	books := []*Book{
		{Title: "Go开发"},
		{Title: "前端开发"},
		{Title: "XX开发"},
	}

	iterator := NewBookIterator(books)
	IteratorFunc(iterator)
}
