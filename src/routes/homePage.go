package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type code struct {
	Code string
}

func GetCode(e echo.Context) error {
	return e.Redirect(http.StatusMovedPermanently, "https://github.com/paij0se/api-deno-compiler#this-a-simple-api-that-execute-your-deno-code-and-send-you-the-output-has-not-limit-per-request")
}
