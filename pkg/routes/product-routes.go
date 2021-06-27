package routes

import (
	. "github.com/KaT0819/sample-go-echo-api/pkg/controllers"
	"github.com/labstack/echo/v4"
)

var ProductsRoutes = func(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/products", GetProducts)
	e.GET("/products/:id", GetProductById)
	e.POST("/products", SaveProduct)
	e.PUT("/products/:id", UpdateProduct)
	e.DELETE("/products/:id", DeleteProduct)
}
