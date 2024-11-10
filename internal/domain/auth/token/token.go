package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/gennadyterekhov/auth-microservice/internal/models/jwtclaims"
	"github.com/golang-jwt/jwt/v5"

	"github.com/pkg/errors"
)

func CreateToken(user *models.User) (string, error) {
	var (
		token         *jwt.Token
		tokenAsString string
		err           error
	)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(user))

	tokenAsString, err = token.SignedString(getJwtSigningKey())
	if err != nil {
		return "", err
	}

	return tokenAsString, nil
}

func newClaims(user *models.User) *jwtclaims.Claims {
	return &jwtclaims.Claims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().AddDate(1, 0, 0)},
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		NotBefore: &jwt.NumericDate{Time: time.Now()},
		Issuer:    "server",
		Subject:   user.Login,
		Audience:  jwt.ClaimStrings{},
		UserID:    user.ID,
	}
}

func getJwtSigningKey() []byte {
	fromEnv, ok := os.LookupEnv("JWT_SIGNING_KEY")
	if ok {
		return []byte(fromEnv)
	}

	return []byte("")
}

func ValidateToken(token string, login string) error {
	claims, err := getClaimsFromToken(token)
	if err != nil {
		return err
	}
	sub, err := claims.GetSubject()
	if err != nil {
		return err
	}

	if sub == login {
		return nil
	}

	return fmt.Errorf("token did not authenticate selected user")
}

func getClaimsFromToken(token string) (*jwtclaims.Claims, error) {
	claims := &jwtclaims.Claims{}

	pureToken := GetPureTokenFromHeaderValue(token)

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	_, err := jwt.ParseWithClaims(
		pureToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return getJwtSigningKey(), nil
		},
	)
	if err != nil {
		logger.Errorln("could not parse token ", pureToken)
		return nil, errors.Wrap(err, "error when parsing token")
	}

	return claims, nil
}

func GetPureTokenFromHeaderValue(header string) string {
	return strings.Replace(header, "Bearer ", "", 1)
}

func GetIDAndLoginFromToken(token string) (int64, string, error) {
	claims, err := getClaimsFromToken(token)
	if err != nil {
		return 0, "", err
	}

	login, err := claims.GetSubject()
	if err != nil {
		return 0, "", err
	}
	id, err := claims.GetUserID()
	if err != nil {
		return 0, "", err
	}

	return id, login, nil
}
