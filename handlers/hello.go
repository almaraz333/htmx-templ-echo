package handlers

import (
	"html-templ-echo/templates/hello"
	"html-templ-echo/types"
	"html-templ-echo/utils"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h HelloHandler) HandleHelloShow(c echo.Context) error {
	helloStruct := types.Hello{
		Name: "Ranger Boy",
	}
	return utils.Render(c, hello.Show(helloStruct))
}
