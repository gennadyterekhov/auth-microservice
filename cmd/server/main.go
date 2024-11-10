package main

import (
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/app"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

func main() {
	fmt.Println("server initialization")

	appInstance := app.New()

	fmt.Println("server initialized successfully")

	err := appInstance.StartServer()
	if err != nil {
		logger.Errorln(err.Error())
		panic(err)
	}
}
