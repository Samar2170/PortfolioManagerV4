package mw

import (
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/samar2170/portfolio-manager-v4/pkg/response"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	UserCID  string `json:"user_cid"`
	jwt.StandardClaims
}

func JwtMiddleware(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println(c.Request().Header)

			if c.Request().Method == "OPTIONS" {
				return next(c)
			}
			if c.Request().URL.Path == "/api/v1/signup" || c.Request().URL.Path == "/api/v1/login" {
				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(response.UnauthorizedResponseEcho("Missing Authorization Header"))
			}
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
			token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				return c.JSON(response.UnauthorizedResponseEcho(err.Error()))
			}
			claims, ok := token.Claims.(*JwtCustomClaims)
			if !ok {
				return c.JSON(response.UnauthorizedResponseEcho("Invalid Token"))
			}
			log.Println(claims)
			c.Set("user_cid", claims.UserCID)
			c.Set("username", claims.Username)
			return next(c)
		}
	}
}
