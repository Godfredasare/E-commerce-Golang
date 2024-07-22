package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRETE_KEY = []byte(os.Getenv("SECRETE_KEY"))

func CreateToken(id, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": id,
			"email":  email,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(SECRETE_KEY)

}

func VerifyToken(tokenString string) (string, error) {
	parseToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpect sign method")
		}
		return SECRETE_KEY, nil
	})

	if err != nil {
		return "", errors.New("unexpect sign method")
	}

	if !parseToken.Valid {
		return "", errors.New("invalid parse token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userId := claims["userId"].(string)

	return userId, nil

}
