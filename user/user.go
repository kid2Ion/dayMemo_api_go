package user

import (
	"net/http"
	"os"

	"github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	gstorage "github.com/hiroki-kondo-git/dayMemo_api_go/gstorage"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	myutil "github.com/hiroki-kondo-git/dayMemo_api_go/utility"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func Signup(ctx echo.Context) error {
	user := new(model.Users)

	// リクエストボディからuser情報取得
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

	iconbase64 := user.IconBase64
	backetName := os.Getenv("BACKET_USER")
	iconName, err := myutil.RandomString(10)
	if err != nil {
		return err
	}

	if err := gstorage.UploadFile(backetName, iconName, iconbase64); err != nil {
		return err
	}
	user.IconBase64 = ""
	user.IconUrl = "https://storage.googleapis.com/daymemo-user/" + iconName

	model.CreateUser(user)
	user.Password = ""

	return ctx.JSON(http.StatusCreated, user)
}

func GetUser(ctx echo.Context) error {
	name := ctx.Param("name")
	u := new(model.Users)

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
	user := new(model.Users)
	if err := ctx.Bind(user); err != nil {
		return err
	}

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	user.ID = uid

	// 既存のuser情報取得
	thisUser := model.FindUserByUid(user)
	if thisUser.ID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	}

	// nameのvalidation
	if u := model.FindUser(user); u.ID != "" && u.UserName != thisUser.UserName {
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

	// gcsimgUpdate
	if user.IconBase64 != "" {
		backetName := os.Getenv("BACKET_USER")
		iconName := thisUser.IconUrl[44:]
		if err := gstorage.DeleteFile(backetName, iconName); err != nil {
			return err
		}

		iconbase64 := user.IconBase64
		iconName, err := myutil.RandomString(10)
		if err != nil {
			return err
		}

		if err := gstorage.UploadFile(backetName, iconName, iconbase64); err != nil {
			return err
		}

		user.IconUrl = "https://storage.googleapis.com/daymemo-user/" + iconName
	} else {
		user.IconUrl = thisUser.IconUrl
	}

	user.IconBase64 = ""
	user = model.UpdateUser(user)

	return ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx echo.Context) error {
	user := new(model.Users)
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
