package main

import (
	"fmt"

	"github.com/KaT0819/sample-go-echo-api/pkg/config"
	"github.com/KaT0819/sample-go-echo-api/pkg/controllers"
	"github.com/KaT0819/sample-go-echo-api/pkg/middlewares"
	"github.com/KaT0819/sample-go-echo-api/pkg/routes"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	cfg := config.ConfigDatabase{}
	if err := env.Parse(&cfg); err != nil {
		e.Logger.Fatal(err)
	}
	e.Logger.Print(cfg)

	controllers.SetValidator(e)

	// Root level middleware
	middlewares.ProductsMiddlewares(e)

	// Routes
	routes.ProductsRoutes(e)

	// Start server
	e.Logger.Print(fmt.Sprintf("access to http://localhost:%s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
