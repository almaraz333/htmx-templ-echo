package main

import (
	"context"
	"htmx-templ-echo/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Use(withMiddleware)
	app.GET("/", handlers.HandleHelloShow)

	app.Start(":3000")
}

func withMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "name", "Storm")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
