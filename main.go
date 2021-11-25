package main

import (
	"net/http"

	memory "github.com/hiroki-kondo-git/dayMemo_api_go/memories"
	user "github.com/hiroki-kondo-git/dayMemo_api_go/users"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	// memory
	e.POST("memories/new", memory.CreateMemory)
	e.PUT("/memories/:id", memory.UpdateMemory)
	e.GET("/memories/list", memory.GetMemories)
	e.GET("/memories/:id", memory.GetMemory)
	// user
	e.POST("/users/new", user.CreateUser)
	e.PUT("/users/:id", user.UpdateUser)
	e.GET("/users/:id", user.GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(ctx echo.Context) error {
	// myfirebase.InitFirebaseAuth()
	return ctx.String(http.StatusOK, "hello dayMemo")
}
