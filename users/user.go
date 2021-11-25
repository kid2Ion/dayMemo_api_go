package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "successfully create user")
}

func UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	msg := "successfully edit user id:" + id
	return ctx.String(http.StatusOK, msg)
}

func GetUser(ctx echo.Context) error {
	id := ctx.Param("id")
	msg := "successfully get user id:" + id
	return ctx.String(http.StatusOK, msg)
}
