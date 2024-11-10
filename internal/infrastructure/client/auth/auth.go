package auth

import (
	"context"
	"os"
	"path"

	swaggerClient "github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger"
	"github.com/gennadyterekhov/auth-microservice/internal/project"
	"golang.org/x/oauth2"
)

func getTokenFilename() (string, error) {
	pr, err := project.GetProjectRoot()
	if err != nil {
		return "", err
	}
	return path.Join(pr, "cmd/client/token"), nil
}

func SetToken(token string) error {
	tf, err := getTokenFilename()
	if err != nil {
		return err
	}
	return os.WriteFile(tf, []byte(token), 0o777)
}

// GetToken gets token from file near the binary
func GetToken() (string, error) {
	tf, err := getTokenFilename()
	if err != nil {
		return "", err
	}

	fileBytes, err := os.ReadFile(tf)
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}

// GetContext gets context with token
func GetContext() (context.Context, error) {
	token, err := GetToken()
	if err != nil {
		return context.Background(), err
	}

	tokenObject := &oauth2.Token{
		AccessToken:  token,
		RefreshToken: token,
	}

	tokenSource := oauth2.StaticTokenSource(tokenObject)
	return context.WithValue(context.Background(), swaggerClient.ContextOAuth2, tokenSource), nil
}
