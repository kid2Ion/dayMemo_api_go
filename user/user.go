package user

import (
	"fmt"
	"net/http"

	"github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
)

// dbにuser登録してuser情報(password以外)返す
func Signup(ctx echo.Context) error {
	user := new(model.User)

	// リクエストボディからuser情報取得
	if err := ctx.Bind(user); err != nil {
		return err
	}

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	user.ID = uid

	if user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid or password",
		}
	}

	if u := model.FindUser(&model.User{Name: user.Name}); u.ID != "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name already exists",
		}
	}

	model.CreateUser(user)
	user.Password = ""

	return ctx.JSON(http.StatusCreated, user)
}

// func UpdateUser(ctx echo.Context) error {
// 	id := ctx.Param("id")
// 	msg := "successfully edit user id:" + id
// 	return ctx.String(http.StatusOK, msg)
// }

// func GetUser(ctx echo.Context) error {
// 	id := ctx.Param("id")
// 	msg := "successfully get user id:" + id
// 	return ctx.String(http.StatusOK, msg)
// }
