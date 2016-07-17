//handler GET request

package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

//define method which will handle GET request
func Get(c echo.Context) error {
	return c.JSON(http.StatusOK, http.StatusOK)
}
