package inits

import (
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
)

func InitFactorySuite[T interfaces.WithFactory](genericSuite T) {
	fmt.Println("InitFactorySuite ")
	fmt.Println()
	InitDbSuite(genericSuite)
	genericSuite.SetFactory(factories.NewFactory(genericSuite.GetRepository()))
}
