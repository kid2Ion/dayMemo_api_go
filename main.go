package main

import (
	"net/http"

	myfirebase "github.com/hiroki-kondo-git/dayMemo_api_go/firebase"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(ctx echo.Context) error {
	myfirebase.InitFirebaseAuth()
	return ctx.String(http.StatusOK, "hello dayMemo")
}
