package output

import (
	"fmt"
)

type Mock struct {
	vals []any
}

func NewMock() *Mock {
	return &Mock{
		vals: make([]any, 0),
	}
}

func (m *Mock) Println(a ...any) {
	m.vals = append(m.vals, a...)
	fmt.Println(a...)
}

func (m *Mock) Get() []any {
	return m.vals
}
