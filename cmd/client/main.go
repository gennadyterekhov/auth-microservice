package main

import (
	"os"

	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/category/create"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/category/read"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/login"
	createOrder "github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/order/create"
	delete2 "github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/order/delete"
	read2 "github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/order/read"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/order/search"
	updateOrder "github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/order/update"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/register"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/root"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/user"
)

func main() {
	rootCmd := root.New()
	loginCmd := login.New()
	registerCmd := register.New()
	userCmd := user.New()
	getCategoriesCmd := read.New()
	createCategoryCmd := create.New()

	createOrderCmd := createOrder.New()
	readOrderCmd := read2.New()
	readListOrderCmd := read2.NewList()
	updateOrderCmd := updateOrder.New()
	deleteOrderCmd := delete2.New()
	buyOrderCmd := delete2.NewBuy()
	searchOrdersCmd := search.New()
	relevantOrdersCmd := search.NewRelevant()

	rootCmd.AddCommand(
		loginCmd,
		registerCmd,
		userCmd,
		getCategoriesCmd,
		createCategoryCmd,
		createOrderCmd,
		readOrderCmd,
		updateOrderCmd,
		deleteOrderCmd,
		buyOrderCmd,
		readListOrderCmd,
		readOrderCmd,
		searchOrdersCmd,
		relevantOrdersCmd,
	)

	// Execute adds all child commands to the root command and sets flags appropriately.
	// This is called by main.main(). It only needs to happen once to the rootCmd.
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
