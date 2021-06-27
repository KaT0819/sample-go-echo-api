package main

import (
	"fmt"
	"os"

	"github.com/KaT0819/sample-go-echo-api/pkg/controllers"
	"github.com/KaT0819/sample-go-echo-api/pkg/middlewares"
	"github.com/KaT0819/sample-go-echo-api/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "1323"
	}

	// Echo instance
	e := echo.New()
	controllers.SetValidator(e)

	// Root level middleware
	middlewares.ProductsMiddlewares(e)

	// Routes
	routes.ProductsRoutes(e)

	// Start server
	e.Logger.Print(fmt.Sprintf("access to http://localhost:%s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
