package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// JwtFilter - verify token jwt
func JwtFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString != "" {
			token, err := validateToken(tokenString)
			if !token.Valid {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			fmt.Println("Authorization Bearer header required")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func extractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	return token, err
}

// GetLoggedUserID - extract logged user id from jwt
func GetLoggedUserID(c *gin.Context) (uuid.UUID, error) {
	tokenString := extractToken(c)
	if tokenString != "" {
		token, err := validateToken(tokenString)
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				uid := fmt.Sprint(claims["sub"])
				if uid == "" {
					return uuid.Nil, nil
				}
				return uuid.FromString(uid)
			}
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		fmt.Println("Authorization Bearer header required")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return uuid.Nil, nil
}
