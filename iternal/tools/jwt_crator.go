package tools

import (
	"Graduation_Project/config"
	"Graduation_Project/iternal/domain"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func JWTCreator(user domain.Register) (string, error, int) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Timur",
			Subject:   "Authorization",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user.ID,
	})

	strToken, err := token.SignedString([]byte(config.Env.SecretKey))

	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	return strToken, nil, http.StatusOK
}
