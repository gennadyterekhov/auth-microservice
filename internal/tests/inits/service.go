package inits

import (
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
)

func InitServiceSuite[T interfaces.WithService](genericSuite T, srv any) {
	fmt.Println("InitServiceSuite ")
	fmt.Println()

	InitFactorySuite(genericSuite)
	genericSuite.SetService(srv)
}
