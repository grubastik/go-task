
package server

import (
    "net/http"
    "fmt"
    "github.com/grubastik/restApi/handler"
	"github.com/grubastik/restApi/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
)

type (
    Server struct {
        Echo *echo.Echo
    }
)

func New() (s *Server) {
    s = &Server{Echo: echo.New()}
    return
}

func (server *Server)RunServer() {

    server.Echo.SetHTTPErrorHandler(MyServerHTTPErrorHandler)

	server.Echo.Use(echomiddleware.Logger())
	server.Echo.Use(echomiddleware.Recover())

    // Login route
	server.Echo.POST("/login", handler.Login)

    r := server.Echo.Group("/api")

	// Middleware
	r.Use(middleware.JWT([]byte("test")))

	// Routes
	r.POST("", handler.Post)
	r.GET("", handler.Get)
	r.POST("/", handler.Post)
	r.GET("/", handler.Get)
	r.GET("/:id", handler.Get)
	r.PUT("/:id", handler.Put)
	r.DELETE("/:id", handler.Delete)

	// Start server
	server.Echo.Run(standard.New(":8080"))
}

func MyServerHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
//	msg := http.StatusText(code)
//	var he echo.HTTPError
	he, ok := err.(*echo.HTTPError);
	fmt.Printf("%v\n",he);
	if ok {
		code = he.Code
//		msg = he.Message
	}
	if !c.Response().Committed() {
		if c.Request().Method() == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			c.JSON(code, he)
//			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScriptCharsetUTF8)
//	        c.Response().WriteHeader(code)
		}
	}
}

