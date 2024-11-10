package login

import (
	"context"
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger/models"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/input"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "login",
		Long:  `login`,
		Run:   Run,
	}
}

func Run(cmd *cobra.Command, args []string) {
	conf := swagger.NewConfiguration(false) // TODO add https
	client := swagger.NewAPIClient(conf)

	inp := input.New()
	run(client, inp)
}

func run(client *swagger.APIClient, inp input.Interface) {
	ctx := context.Background()

	var err error
	var login string
	var password string

	fmt.Println("login: ")

	err = inp.ScanStrings(&login)
	if err != nil {
		logger.Errorln(err.Error())
		return
	}
	fmt.Println("password: ")

	err = inp.ScanStrings(&password)
	if err != nil {
		logger.Errorln(err.Error())
		return
	}

	grpcResp, _, err := client.AuthApi.ArtDealersLogin(ctx, models.ProtobufLoginRequest{Login: login, Password: password})
	if err != nil {
		logger.Errorln(err.Error())
		return
	}
	fmt.Println("token:", grpcResp.Token)
	err = auth.SetToken(grpcResp.Token)
	if err != nil {
		logger.Errorln(err.Error())
		return
	}
}
