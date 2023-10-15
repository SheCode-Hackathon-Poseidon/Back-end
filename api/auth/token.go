package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken is...
func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// TokenValid is...
func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return err
	}

	switch claims := token.Claims.(type) {
	case jwt.MapClaims:
		if token.Valid {
			Pretty(claims)
		} else {
			return nil
		}
	default:
		return nil
	}

	return nil
}

// ExtractToken is...
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()

	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}



// Pretty is...
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Println(err)

		return
	}

	fmt.Println(string(b))
}
