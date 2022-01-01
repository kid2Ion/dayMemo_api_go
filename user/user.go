package user

import (
	"net/http"
	"time"

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

	if u := model.FindUser(&model.User{UserName: user.UserName}); u.ID != "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name already exists",
		}
	}

	model.CreateUser(user)
	user.Password = ""
	user.Email = ""
	user.CreatedAt = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)

	return ctx.JSON(http.StatusCreated, user)
}

func GetUser(ctx echo.Context) error {
	name := ctx.Param("name")
	u := new(model.User)

	// 懸念：適当にnameいれてrequest投げてそれが当たったら、uidがresponseに入る→いいのか？
	u.UserName = name
	user := model.FindUser(&model.User{UserName: u.UserName})
	if user.ID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}
	// resにいらないものはdefault値
	user.Email = ""
	user.Password = ""
	user.CreatedAt = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
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

	// 一応存在しなければerrで返す（firebaseからとってくるから存在しないことはない）
	u := model.FindUser(&model.User{ID: user.ID})
	if u.ID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "failed to update: user not found",
		}
	}

	user = model.UpdateUser(user)
	user.ID = ""

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
