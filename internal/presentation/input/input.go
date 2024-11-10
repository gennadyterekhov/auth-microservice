package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Interface interface {
	ScanInts(...*int64) error
	ScanStrings(...*string) error
}

var (
	_ Interface = New()
	_ Interface = NewMock()
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s Service) ScanInts(a ...*int64) error {
	var err error
	for i := range a {
		_, err = fmt.Scan(a[i])
		if err != nil {
			return err
		}
	}

	return err
}

func (s Service) ScanStrings(a ...*string) error {
	var err error

	in := bufio.NewReader(os.Stdin)

	for i := range a {

		line, err := in.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.Trim(line, "\n ")

		*a[i] = line
	}

	return err
}
