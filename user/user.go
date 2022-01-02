package user

import (
	"net/http"

	"github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
)

func Signup(ctx echo.Context) error {
	user := new(model.User)

	// リクエストボディからuser情報取得
	if err := ctx.Bind(user); err != nil {
		return err
	}

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	user.ID = uid

	if user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid or password",
		}
	}

	if u := model.FindUser(user); u.ID != "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name already exists",
		}
	}

	model.CreateUser(user)
	user.Password = ""

	return ctx.JSON(http.StatusCreated, user)
}

func GetUser(ctx echo.Context) error {
	name := ctx.Param("name")
	u := new(model.User)

	u.UserName = name
	user := model.FindUser(u)
	if user.ID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	user.Password = ""
	return ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx echo.Context) error {
	user := new(model.User)
	if err := ctx.Bind(user); err != nil {
		return err
	}

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	user.ID = uid

	user = model.UpdateUser(user)

	return ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx echo.Context) error {
	user := new(model.User)
	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	user.ID = uid

	model.DeleteUser(user)
	return ctx.JSON(http.StatusOK, "successfully user delete")
}
