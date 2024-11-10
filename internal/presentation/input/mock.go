package input

import (
	"fmt"
	"sync/atomic"
)

type Mock struct {
	strs       []string
	ints       []int64
	nextInt    atomic.Int64
	nextString atomic.Int64
}

func NewMock() *Mock {
	return &Mock{
		strs: make([]string, 0),
		ints: make([]int64, 0),
	}
}

func (m *Mock) AddStrings(a ...string) {
	m.strs = append(m.strs, a...)
}

func (m *Mock) AddInts(i ...int64) {
	m.ints = append(m.ints, i...)
}

func (m *Mock) scanInt(a *int64) error {
	i := 0

	if m.nextInt.Add(int64(i)) >= int64(len(m.ints)) {
		return fmt.Errorf("no item in buffer")
	}

	ind := int64(m.nextInt.Add(int64(i)))

	*a = m.ints[ind]
	m.nextInt.Store(m.nextInt.Add(1))

	return nil
}

func (m *Mock) ScanInts(a ...*int64) error {
	var err error
	for i := range a {
		err = m.scanInt(a[i])
		if err != nil {
			return err
		}
	}

	return err
}

func (m *Mock) scanString(a *string) error {
	i := 0

	if m.nextString.Add(int64(i)) >= int64(len(m.strs)) {
		return fmt.Errorf("no item in buffer")
	}

	ind := int64(m.nextString.Add(int64(i)))

	*a = m.strs[ind]
	m.nextString.Store(m.nextString.Add(1))

	return nil
}

func (m *Mock) ScanStrings(a ...*string) error {
	var err error
	for i := range a {
		err = m.scanString(a[i])
		if err != nil {
			return err
		}
	}

	return err
}
