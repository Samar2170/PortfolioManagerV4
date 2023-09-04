package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samar2170/portfolio-manager-v4/internal"
	"github.com/samar2170/portfolio-manager-v4/internal/bulkupload"
	"github.com/samar2170/portfolio-manager-v4/pkg/mw"
	"github.com/samar2170/portfolio-manager-v4/pkg/response"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mw.JwtMiddleware(SigningKey))
	subroute := e.Group("/api/v1")
	subroute.POST("/signup", signup)
	subroute.POST("/login", login)

	subroute.POST("/register-account/:accountType", registerAccount)
	e.Logger.Fatal(e.Start(":8080"))

}

func signup(c echo.Context) error {
	signupRequest := new(internal.SignupRequest)

	if err := c.Bind(signupRequest); err != nil {
		return c.JSON(response.BadRequestResponseEcho(err.Error()))
	}
	err := internal.Signup(*signupRequest)
	if err != nil {
		return c.JSON(response.BadRequestResponseEcho(err.Error()))
	}
	return c.JSON(response.SuccessResponseEcho("User Created Successfully"))
}

func login(c echo.Context) error {
	loginRequest := new(internal.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(response.BadRequestResponseEcho(err.Error()))
	}
	if loginRequest.Username == "" || loginRequest.Password == "" {
		return c.JSON(response.BadRequestResponseEcho("Username or Password cannot be empty"))
	}
	token, err := internal.Login(*loginRequest)
	if err != nil {
		return c.JSON(response.BadRequestResponseEcho(err.Error()))
	}
	return c.JSON(response.SuccessResponseEcho(token))
}

func registerAccount(c echo.Context) error {
	accountType := c.Param("accountType")
	userCID := c.Get("user_cid").(string)
	username := c.Get("username").(string)
	log.Println("user: ", username, "user_cid: ", userCID)
	switch accountType {
	case "bank":
		bankAccountRequest := new(internal.BankAccountRequest)
		if err := c.Bind(bankAccountRequest); err != nil {
			return c.JSON(response.BadRequestResponseEcho(err.Error()))
		}
		err := internal.RegisterBankAccount(*bankAccountRequest, userCID)
		if err != nil {
			return c.JSON(response.BadRequestResponseEcho(err.Error()))
		}
		return c.JSON(response.SuccessResponseEcho("Bank Account Registered Successfully"))
	case "demat":
		dematAccountRequest := new(internal.DematAccountRequest)
		if err := c.Bind(dematAccountRequest); err != nil {
			return c.JSON(response.BadRequestResponseEcho(err.Error()))
		}
		err := internal.RegisterDematAccount(*dematAccountRequest, userCID)
		if err != nil {
			return c.JSON(response.BadRequestResponseEcho(err.Error()))
		}
		return c.JSON(response.SuccessResponseEcho("Demat Account Registered Successfully"))
	default:
		return c.JSON(response.BadRequestResponseEcho("Invalid Account Type"))
	}
}

func downloadTradeTemplate(c echo.Context) error {
	return c.Attachment("assets/trade-template.xlsx", "Trade Template")
}

func uploadTradeSheet(c echo.Context) error {
	userCID := c.Get("user_cid").(string)
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(response.BadRequestResponseEcho(err.Error()))
	}
	err = bulkupload.SaveBulkUploadFile(file, userCID)
	if err != nil {
		return c.JSON(response.BadRequestResponseEcho(err.Error()))
	}
	return c.JSON(response.SuccessResponseEcho("Successfully Upload"))
}
