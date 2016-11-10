//handler GET request

package handler

import (
    "net/http"
	"time"

    "github.com/labstack/echo"
  	"github.com/dgrijalva/jwt-go"
)

//define method which will handle GET request
func Login(c echo.Context) error {
    if c.FormValue("user") == "test" && c.FormValue("password") == "test" {
        // Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = c.FormValue("user")
		claims["exp"] = time.Now().Add(time.Hour).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("test"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
