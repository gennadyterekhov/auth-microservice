package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gennadyterekhov/auth-microservice/internal/app"
	"github.com/gennadyterekhov/auth-microservice/internal/project"
	"github.com/gennadyterekhov/auth-microservice/internal/project/config"
)

func main() {
	certFilename, keyFilename, serverConfig, appInstance, err := getDeps()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("listening on https://" + serverConfig.Addr)
	err = http.ListenAndServeTLS(serverConfig.Addr, certFilename, keyFilename, appInstance.Router())
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getDeps() (string, string, *config.Config, *app.App, error) {
	certFilename, keyFilename, err := getTlsFilenames()
	if err != nil {
		return "", "", nil, nil, err
	}

	serverConfig, appInstance, err := getAppInstance()
	if err != nil {
		return "", "", nil, nil, err
	}
	return certFilename, keyFilename, serverConfig, appInstance, nil
}

func getTlsFilenames() (string, string, error) {
	pr, err := project.GetProjectRoot()
	if err != nil {
		return "", "", err
	}

	return path.Join(pr, "certificates", "server.crt"), path.Join(pr, "certificates", "server.key"), nil
}

func getAppInstance() (*config.Config, *app.App, error) {
	serverConfig, err := config.New()
	if err != nil {
		return nil, nil, err
	}

	appInstance, err := app.New(serverConfig.DBDsn)
	return serverConfig, appInstance, err
}
