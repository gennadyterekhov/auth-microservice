package output

import (
	"fmt"
)

type Interface interface {
	Println(...any)
}

var (
	_ Interface = New()
	_ Interface = NewMock()
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s Service) Println(a ...any) {
	fmt.Println(a...)
}
