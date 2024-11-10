package middleware

import (
	"context"
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/consts"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/token"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")

		pureToken := token.GetPureTokenFromHeaderValue(authHeader)

		var id int64
		id, _, err := token.GetIDAndLoginFromToken(pureToken)
		if err != nil {
			logger.Errorln(err.Error())
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), consts.ContextUserIDKey, id)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
