package tools

import (
	"Graduation_Project/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func JwtParsUserID(c *gin.Context) uuid.UUID {
	claims := &Claims{}
	userId := uuid.UUID{}

	strToken := c.Request.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(strToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Env.SecretKey), nil
	})
	if err != nil {
		CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "ReadAll").Warn(err)
		return uuid.UUID{}
	}
	if token.Valid {
		userId = claims.UserID
	}
	return userId
}
