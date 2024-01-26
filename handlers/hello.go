package handlers

import (
	"htmx-templ-echo/templates/hello"
	"htmx-templ-echo/types"
	"htmx-templ-echo/utils"

	"github.com/labstack/echo/v4"
)

func HandleHelloShow(c echo.Context) error {
	helloStruct := types.Hello{
		Name: "Ranger Boy",
	}
	return utils.Render(c, hello.Show(helloStruct))
}
