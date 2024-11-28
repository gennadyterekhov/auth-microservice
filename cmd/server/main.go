package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/app"
	"github.com/gennadyterekhov/auth-microservice/internal/project/config"
)

func main() {
	fmt.Println("server initialization")
	serverConfig, appInstance, err := getAppInstance()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("got app instance")

	fmt.Println("server initialized successfully")

	fmt.Println("listening with https on " + serverConfig.Addr)

	err = http.ListenAndServeTLS(serverConfig.Addr, "cmd/server/server.crt", "cmd/server/server.key", appInstance.Router())
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getAppInstance() (*config.Config, *app.App, error) {
	serverConfig, err := config.New()
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("got server config")

	appInstance, err := app.New(serverConfig.DBDsn)
	return serverConfig, appInstance, err
}
