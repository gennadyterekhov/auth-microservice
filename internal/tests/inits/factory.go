package inits

import (
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
)

func InitFactorySuite[T interfaces.WithFactory](genericSuite T) {
	InitDbSuite(genericSuite)
	genericSuite.SetFactory(factories.NewFactory(genericSuite.GetRepository()))
}
