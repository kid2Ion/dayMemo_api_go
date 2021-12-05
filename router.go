package main

import (
	"net/http"

	auth "github.com/hiroki-kondo-git/dayMemo_api_go/auth"
	memory "github.com/hiroki-kondo-git/dayMemo_api_go/memories"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", hello)

	// user
	e.POST("/users/new", auth.Signup)
	// e.PUT("/users/:id", user.UpdateUser)
	// e.GET("/users/:id", user.GetUser)

	api := e.Group("/api")
	// api.Use(middleware.JWTWithConfig(auth.Config))
	// memory
	api.POST("memories/new", memory.CreateMemory)
	// e.PUT("/memories/:id", memory.UpdateMemory)
	api.GET("/memory/list", memory.GetMemories)
	// api.GET("/memories/:id", memory.GetMemory)
	api.DELETE("/memory/:id", memory.DeleteMemory)

	return e
}

func hello(ctx echo.Context) error {
	// myfirebase.InitFirebaseAuth()
	return ctx.String(http.StatusOK, "hello dayMemo")
}
