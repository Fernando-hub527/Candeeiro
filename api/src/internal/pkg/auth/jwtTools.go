package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Minute * 5).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString("secretKey")
}

func ValidateToken(r *http.Request) (string, error) {
	tokenString := strings.Replace(r.Header.Get("Authorization"), "bearer ", "", 0)

	token, err := jwt.Parse(tokenString, returnKeyVerifyToken)

	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := permissions["userId"].(string)
		return userId, nil
	}

	return "", errors.New("invalid token")
}

func returnKeyVerifyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signing method unexpected! %v", token.Header["alg"])
	}

	return "config.SecretKey", nil
}
