package memory

import (
	"net/http"
	"strconv"

	"github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	"github.com/hiroki-kondo-git/dayMemo_api_go/model"
	"github.com/labstack/echo"
)

func CreateMemory(ctx echo.Context) error {
	memory := new(model.Memory)
	// リクエストボディからuser情報取得
	if err := ctx.Bind(memory); err != nil {
		return err
	}

	if memory.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or title",
		}
	}

	// token→memory{UID}=user{ID}
	uid := auth.UserIDFromToken(ctx)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	memory.UID = uid
	model.CreateMemory(memory)

	return ctx.JSON(http.StatusOK, memory)
}

func GetMemories(ctx echo.Context) error {
	uid := auth.UserIDFromToken(ctx)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
	// todo 月ごとのmemory取得の実装
	// year_month := ctx.QueryParam("year_month")
	// msg := "successfully get memories list year_month:" + year_month

	memoryList := model.FindMemories(&model.Memory{UID: uid})
	return ctx.JSON(http.StatusOK, memoryList)
}

// todo update処理
// func UpdateMemory(ctx echo.Context) error {
// 	uid := auth.UserIDFromToken(ctx)
// 	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
// 		return echo.ErrNotFound
// 	}

// 	memoryID, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		return echo.ErrNotFound
// 	}

// 	memories := model.FindMemories(&model.Memory{ID: memoryID, UID: uid})
// 	if len(memories) == 0 {
// 		return echo.ErrNotFound
// 	}
// memory := memories[0]
// }

func DeleteMemory(ctx echo.Context) error {
	uid := auth.UserIDFromToken(ctx)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	memoryID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteMemory(&model.Memory{ID: memoryID, UID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return ctx.NoContent(http.StatusOK)
}
