package mw

import "github.com/labstack/echo/v4"

type CookieUser struct {
	Token string
}

func TemplateLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("username")
		if err == nil && cookie.Value != "" {
			user := &CookieUser{Token: cookie.Value}
			ctx.Set("user", user)
		}
		return next(ctx)
	}
}
