package register

import (
	"context"
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger/models"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "register",
		Long:  `register`,
		Run:   Run,
	}
}

func Run(cmd *cobra.Command, args []string) {
	conf := swagger.NewConfiguration(false) // TODO add https
	client := swagger.NewAPIClient(conf)

	// TODO make beautiful common output
	fmt.Println("register command")
	ctx := context.Background()

	// TODO make input loop
	var err error
	var login string
	var password string

	fmt.Println("login: ")

	_, err = fmt.Scan(&login)
	if err != nil {
		logger.Errorln(err.Error())
		return
	}
	fmt.Println("password: ")

	_, err = fmt.Scan(&password)
	if err != nil {
		logger.Errorln(err.Error())
		return
	}

	// grpcResp, httpResp, err
	grpcResp, _, err := client.AuthApi.ArtDealersRegister(ctx, models.ProtobufRegisterRequest{Login: login, Password: password})
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
