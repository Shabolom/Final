package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CookieCreate(c *gin.Context, token string) {
	cookie := http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 3),
	}

	http.SetCookie(c.Writer, &cookie)
}
