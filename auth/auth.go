package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

// dbにuser登録してuser情報(password以外)返す
func Signup(ctx echo.Context) error {
	user := new(model.User)

	// リクエストボディからuser情報取得
	if err := ctx.Bind(user); err != nil {
		return err
	}

	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	if u := model.FindUser(&model.User{Name: user.Name}); u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name already exists",
		}
	}

	model.CreateUser(user)
	user.Password = ""

	return ctx.JSON(http.StatusCreated, user)
}

// ログインしてtoken返す
func Login(ctx echo.Context) error {
	u := new(model.User)

	// リクエストボディからuser情報取得
	if err := ctx.Bind(u); err != nil {
		return err
	}

	user := model.FindUser(&model.User{Name: u.Name})
	if user.ID == 0 || user.Password != u.Password {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// トークンからuidを取得して返す
func UserIDFromToken(ctx echo.Context) int {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
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
