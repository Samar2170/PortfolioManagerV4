package main

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samar2170/portfolio-manager-v4/internal"
	"github.com/samar2170/portfolio-manager-v4/pkg/mw"
)

var cookieStore = sessions.NewCookieStore([]byte("portfolio-manager"))

func loginPage(c echo.Context) error {
	tpl, err := template.ParseFiles("template/login.html")
	if err != nil {
		c.Logger().Error("parse file error:", err)
		return err
	}
	data := map[string]interface{}{
		"msg": c.QueryParam("msg"),
	}

	if user, ok := c.Get("user").(*mw.CookieUser); ok {
		data["token"] = user.Token
	} else {
		sess := getCookieSession(c)
		if flashes := sess.Flashes("username"); len(flashes) > 0 {
			data["username"] = flashes[0]
		}
		sess.Save(c.Request(), c.Response())
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, buf.String())
}

func login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	passwd := ctx.FormValue("password")
	if username != "" && passwd != "" {
		loginRequest := new(internal.LoginRequest)
		loginRequest.Password = passwd
		loginRequest.Username = username
		token, err := internal.Login(*loginRequest)
		if err != nil {
			return err
		}
		cookie := &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
		}
		ctx.SetCookie(cookie)
		return ctx.Redirect(http.StatusSeeOther, "/")

	}
	return ctx.Redirect(http.StatusSeeOther, "/?msg=No user nme pass provided")

}

func getCookieSession(ctx echo.Context) *sessions.Session {
	sess, _ := cookieStore.Get(ctx.Request(), "token")
	return sess
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(mw.TemplateLogin)
	e.Use(middleware.Recover())
	subroute := e.Group("/site")

	subroute.POST("/login", login)

	e.Logger.Fatal(e.Start(":2020"))

}
