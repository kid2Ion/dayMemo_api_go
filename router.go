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
	e.POST("/user/new", user.Signup)
	e.GET("/user/:id", user.GetUser) // ここだけ認証なし
	e.PUT("user/update", user.UpdateUser)
	e.DELETE("user/delete", user.DeleteUser)

	// memory
	e.POST("memories/new", memory.CreateMemory)
	e.GET("/memories/list", memory.GetMemoryList)
	e.GET("/memory/:id", memory.GetMemory)
	e.PUT("/memory/update/:id", memory.UpdateMemory)
	e.DELETE("/memory/:id", memory.DeleteMemory)

	return e
}

func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello dayMemo")
}
