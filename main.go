//main entry point of the rest API server
//ideally it should run separate package which will handle all the staff

package main

import (
    "github.com/grubastik/restApi/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/api", handler.Post)
	e.GET("/api", handler.Get)
	e.GET("/api/:id", handler.Get)
	e.PUT("/api/:id", handler.Put)
	e.DELETE("/api/:id", handler.Delete)

	// Start server
	e.Run(standard.New(":8080"))
}
