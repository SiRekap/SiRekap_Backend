package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const hmacSecret = "sirekap-2024"

func Sign(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerString := c.Request.Header.Get("Authorization")
		if bearerString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		split := strings.Split(bearerString, " ")
		tokenString := split[1]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSecret, nil
		})

		if token == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok || token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}
