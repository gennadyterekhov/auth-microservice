package suites

import (
	"fmt"

	"github.com/stretchr/testify/suite"
)

type Abstract struct {
	suite.Suite
}

func (s *Abstract) SetupTest() {
	fmt.Println("(s *Abstract) SetupTest()")
	fmt.Println()
}

func (s *Abstract) TearDownTest() {
}

func (s *Abstract) SetupSuite() {
	fmt.Println("(s *Abstract) SetupSuite()")
	fmt.Println()
}

func (s *Abstract) TearDownSuite() {
	fmt.Println("(s *Abstract) TearDownSuite()")
	fmt.Println()
}
