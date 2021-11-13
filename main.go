package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello dayMemo")
}
