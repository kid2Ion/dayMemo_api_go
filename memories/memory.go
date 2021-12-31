package memory

import (
	"fmt"
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
	// token→memory{UID}=user{ID}
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
	model.CreateMemory(memory)
	return ctx.JSON(http.StatusOK, memory)
}

func GetMemories(ctx echo.Context) error {
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
	uid, err := auth.AuthFirebase(ctx)
	if err != nil {
		return err
	}

	if user := model.FindUser(&model.User{ID: uid}); user.ID == "" {
		return echo.ErrNotFound
	}

	memoryID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Errorf("error get memoryID", err)
	}

	// gorm.modelを展開するためにインスタンス化
	memory := &model.Memory{UID: uid}
	memory.ID = uint(memoryID)
	if err := model.DeleteMemory(memory); err != nil {
		return echo.ErrNotFound
	}

	return ctx.NoContent(http.StatusOK)
}
