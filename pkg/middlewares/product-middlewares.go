package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ProductsMiddlewares = func(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
