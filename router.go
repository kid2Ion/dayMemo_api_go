package main

import (
	"net/http"

	memory "github.com/hiroki-kondo-git/dayMemo_api_go/memories"
	user "github.com/hiroki-kondo-git/dayMemo_api_go/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/", hello)

	// user
	e.POST("/users/new", user.Signup)

	// memory
	e.POST("memories/new", memory.CreateMemory)
	e.GET("/memory/list", memory.GetMemories)
	// api.GET("/memories/:id", memory.GetMemory)
	// e.PUT("/memories/:id", memory.UpdateMemory)
	e.DELETE("/memory/:id", memory.DeleteMemory)

	return e
}

func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello dayMemo")
}
