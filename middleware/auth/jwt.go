package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/chuckkainrath/SensorProject/middleware"
	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("Authorization") // Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			err := errors.NewForbiddenError("Missing auth token")
			middleware.AddResultToContext(r, err, middleware.ErrorKey)
			return
		}
		tknParts := strings.Fields(header)
		if len(tknParts) != 2 || tknParts[0] != "Bearer" {
			err := errors.NewForbiddenError("Invalid token")
			middleware.AddResultToContext(r, err, middleware.ErrorKey)
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(tknParts[1], tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("randomsecretstring"), nil
		})

		if err != nil {
			err := errors.NewForbiddenError("Invalid token")
			middleware.AddResultToContext(r, err, middleware.ErrorKey)
			return
		}

		ctx := r.Context()
		req := r.WithContext(context.WithValue(ctx, middleware.TokenKey, tk))
		*r = *req
		next.ServeHTTP(w, r)
	})
}

func GetTokenData(r *http.Request) interface{} {
	return r.Context().Value(middleware.TokenKey)
}
