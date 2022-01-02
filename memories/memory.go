package memory

import (
	"net/http"
	"strconv"

	auth "github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
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

	// todo validation
	if memory.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "title is empty",
		}
	}

	//todo ここにiconbase64(json)をもとにgstorageあげる処理→imageURLをmemoryに格納して、一緒にcreate
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
	// todo validation
	if err := ctx.Bind(memory); err != nil {
		return err
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
