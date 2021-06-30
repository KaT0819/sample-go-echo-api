package routes

import (
	. "github.com/KaT0819/sample-go-echo-api/pkg/controllers"
	"github.com/KaT0819/sample-go-echo-api/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

var ProductsRoutes = func(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/products", GetProducts, middlewares.ServerMessage2)
	e.GET("/products/:id", GetProductById, middlewares.ServerMessage, middlewares.ServerMessage2)
	e.POST("/products", SaveProduct)
	e.PUT("/products/:id", UpdateProduct)
	e.DELETE("/products/:id", DeleteProduct)
}
