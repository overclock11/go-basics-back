package jwt

import (
	"errors"
	"golangapi/models"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func Process(tk string, JWTSing string) (*models.Claim, bool, string, error) {
	myPassword := []byte(JWTSing)
	var claim models.Claim
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claim, false, "", errors.New("Token invalido, no tiene bearer")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claim, func(token *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})

	if err != nil {
		// revisar contra mongo
	}

	if !tkn.Valid {
		return &claim, false, "", errors.New("Token invalido")
	}

	return &claim, false, "", errors.New("Token invalido")
}
