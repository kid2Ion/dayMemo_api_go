package memory

import (
	"net/http"
	"strconv"
	"time"

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

	if memory.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or title",
		}
	}

	//todo ここにiconbase64(json)をもとにgstorageあげる処理→imageURLをmemoryに格納して、一緒にcreate
	model.CreateMemory(memory)
	memory.CreatedAt = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	memory.UpdatedAt = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)

	return ctx.JSON(http.StatusOK, memory)
}

func GetMemoryList(ctx echo.Context) error {
	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}

	if user := model.FindUser(&model.User{ID: uid}); user.ID == "" {
		return echo.ErrNotFound
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

	memory := model.FindMemory(&model.Memory{ID: m.ID})
	if memory.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "memory not found",
		}
	}
	// 他人がgetできないように
	if memory.UID != uid {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "memory not found,not your memory",
		}
	}

	return ctx.JSON(http.StatusOK, memory)
}

func UpdateMemory(ctx echo.Context) error {
	memory := new(model.Memory)
	if err := ctx.Bind(memory); err != nil {
		return err
	}
	memoryID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	memory.ID = uint(memoryID)

	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}

	memory.UID = uid

	m := model.FindMemory(&model.Memory{ID: uint(memoryID)})
	if m.UID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "memory not found",
		}
	}
	// 他人がupdateできないように
	if m.UID != uid {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "memory not found,not your memory",
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

	memoryID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	memory.ID = uint(memoryID)

	m := model.FindMemory(&model.Memory{ID: uint(memoryID)})
	if m.UID == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "memory not found",
		}
	}
	// 他人がdeleteできないように
	if m.UID != uid {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "memory not found,not your memory",
		}
	}

	model.DeleteMemory(memory)

	return ctx.JSON(http.StatusOK, "successfully delete memory")
}
