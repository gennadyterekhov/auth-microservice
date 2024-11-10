package main

import (
	"os"

	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/login"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/register"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/client/cmds/root"
)

func main() {
	rootCmd := root.New()
	loginCmd := login.New()
	registerCmd := register.New()

	rootCmd.AddCommand(
		loginCmd,
		registerCmd,
	)

	// Execute adds all child commands to the root command and sets flags appropriately.
	// This is called by main.main(). It only needs to happen once to the rootCmd.
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
