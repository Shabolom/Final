package middleware

import (
	"Graduation_Project/config"
	"Graduation_Project/iternal/tools"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// продолжаем работать с хэндлером который идет после мидлвейра  (который был вызван изначально))
		c.Next()

		latency := time.Since(t)
		log.WithField("component", "latency").Info(latency)
	}
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("Authorization") == `` {
			tools.CreateError(http.StatusUnauthorized, errors.New("you're Unauthorized"), c)
			return
		}

		strToken := c.Request.Header.Get("Authorization")

		token, err := jwt.Parse(strToken,
			func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					tools.CreateError(http.StatusUnauthorized, errors.New("you're Unauthorized"), c)
				}
				return []byte(config.Env.SecretKey), nil
			})

		if err != nil {
			tools.CreateError(http.StatusUnauthorized, errors.New("you're Unauthorized"), c)
			return
		}

		if !token.Valid {
			tools.CreateError(http.StatusUnauthorized, errors.New("you're Unauthorized"), c)
			return
		} else {
			c.Next()
		}
	}
}
