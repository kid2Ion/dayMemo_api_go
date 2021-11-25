package memory

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateMemory(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "successfully create memory")
}

func UpdateMemory(ctx echo.Context) error {
	id := ctx.Param("id")
	msg := "successfully edit memory id:" + id
	return ctx.String(http.StatusOK, msg)
}

func GetMemories(ctx echo.Context) error {
	year_month := ctx.QueryParam("year_month")
	msg := "successfully get memories list year_month:" + year_month
	return ctx.String(http.StatusOK, msg)
}

func GetMemory(ctx echo.Context) error {
	id := ctx.Param("id")
	msg := "successfully get memory id:" + id
	return ctx.String(http.StatusOK, msg)
}
