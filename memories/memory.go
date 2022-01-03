package memory

import (
	"net/http"
	"strconv"

	auth "github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	gstorage "github.com/hiroki-kondo-git/dayMemo_api_go/gstorage"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func CreateMemory(ctx echo.Context) error {
	memory := new(model.Memory)
	// リクエストボディからmemory情報取得
	if err := ctx.Bind(memory); err != nil {
		return err
	}

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	memory.UID = uid

	// validation
	validate := validator.New()
	errors := validate.Struct(memory)
	if errors != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: ValidationError(errors),
		}
	}

	imagebase64 := memory.ImageBase64

	gstorage.UploadFile("daymemo-memory", "LGTM.png", imagebase64)
	memory.ImageBase64 = ""
	model.CreateMemory(memory)

	return ctx.JSON(http.StatusOK, memory)
}

func GetMemoryList(ctx echo.Context) error {
	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}

	year := ctx.QueryParam("year")
	month := ctx.QueryParam("month")
	memoryList := model.FindMemories(&model.Memory{UID: uid}, year, month)

	return ctx.JSON(http.StatusOK, memoryList)
}

func GetMemory(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	m := new(model.Memory)
	m.ID = uint(id)
	m.UID = uid

	memory := model.FindMemory(m)
	if memory.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Memory is not found or not your own.",
		}
	}

	return ctx.JSON(http.StatusOK, memory)
}

func UpdateMemory(ctx echo.Context) error {
	memory := new(model.Memory)
	memoryID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	memory.ID = uint(memoryID)

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}

	memory.UID = uid

	m := model.FindMemory(memory)
	if m.UID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Memory is not found or not your own.",
		}
	}

	if err := ctx.Bind(memory); err != nil {
		return err
	}
	// validation
	validate := validator.New()
	errors := validate.Struct(memory)
	if errors != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: ValidationError(errors),
		}
	}

	memory = model.UpdateMemory(memory)

	return ctx.JSON(http.StatusOK, memory)
}

func DeleteMemory(ctx echo.Context) error {
	memory := new(model.Memory)
	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}
	memory.UID = uid

	memoryID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	memory.ID = uint(memoryID)

	m := model.FindMemory(memory)
	if m.UID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Memory is not found or not your own.",
		}
	}

	model.DeleteMemory(memory)

	return ctx.JSON(http.StatusOK, "successfully delete memory")
}

// validationErrMessage
func ValidationError(err error) []string {
	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		var errorMessage string
		fieldName := err.Field()

		switch fieldName {
		case "Title":
			errorMessage = "Title must over 1 less 20 charactors"
		case "Content":
			errorMessage = "Content is required"
		}
		errorMessages = append(errorMessages, errorMessage)
	}
	return errorMessages
}
