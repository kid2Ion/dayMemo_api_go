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
