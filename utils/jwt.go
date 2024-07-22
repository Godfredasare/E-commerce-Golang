package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRETE_KEY = []byte(os.Getenv("SECRETE_KEY"))

func CreateToken(id, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"_id":   id,
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(SECRETE_KEY)

}
