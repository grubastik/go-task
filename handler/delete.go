//handler DELETE request

package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
