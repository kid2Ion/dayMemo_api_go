package user

import (
	"net/http"

	"github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func Signup(ctx echo.Context) error {
	user := new(model.User)

	// リクエストボディからuser情報取得
	// todo validation
	if err := ctx.Bind(user); err != nil {
		return err
	}

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	user.ID = uid

	// validation
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: ValidationError(errors),
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

	if u := model.FindUser(user); u.ID != "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name already exists",
		}
	}

	// validation
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: ValidationError(errors),
		}
	}

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

// validationErrMessage
func ValidationError(err error) []string {
	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		var errorMessage string
		fieldName := err.Field()

		switch fieldName {
		case "UserName":
			var typ = err.Tag()
			switch typ {
			case "min":
				errorMessage = "username must over 4 charactors"
			case "max":
				errorMessage = "username must less 10 charactors"
			case "alphanum":
				errorMessage = "username must use only alpha and number"
			}
		case "DisplayName":
			errorMessage = "displayName must over 1 less 10 charactors"
		case "Password":
			errorMessage = "password must over 6 charactors"
		}
		errorMessages = append(errorMessages, errorMessage)
	}
	return errorMessages
}
