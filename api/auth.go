package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samar2170/portfolio-manager-v4/internal"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/signup", signup)

	e.Logger.Fatal(e.Start(":8080"))

}

func signup(c echo.Context) error {
	signupRequest := new(internal.SignupRequest)

	if err := c.Bind(signupRequest); err != nil {
		return err
	}
	err := internal.Signup(*signupRequest)
	if err != nil {
		return c.String(400,
			err.Error(),
		)
	}
	return c.String(200,
		"User created successfully",
	)
}
