package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"sample/api/auth"
	"sample/api/exitcode"
	"sample/api/responses"
)

// SetMiddlewareJSON is...
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			fmt.Println("It's options")
			allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

			// Handle preflight request
			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			res.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
			res.Header().Set("Access-Control-Expose-Headers", "Authorization")

		} else {
			res.Header().Set("Content-Type", "application/json")

		}

		next(res, req)
	}
}

// SetMiddlewareAuth is...
func SetMiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := auth.TokenValid(req)

		if err != nil {
			responses.ERROR(res, http.StatusUnauthorized, exitcode.UNAUTHORIZED, errors.New("Unauthorized"))

			return
		}

		next(res, req)
	}
}
