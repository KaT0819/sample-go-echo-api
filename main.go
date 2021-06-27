package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

var products = []map[int]string{{1: "mobilies"}, {2: "TV"}, {3: "laptop"}}

// Validator
func (p *ProductValidator) Validate(i interface{}) error {
	if err := p.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Handler
func home(c echo.Context) error {
	return c.HTML(http.StatusOK, `
		<html lang="ja">
		<head><title>Document</title></head>
		<body> Hello Echo </body>
		</html>
	`)
}

func getProductById(c echo.Context) error {
	var product map[int]string
	for _, p := range products {
		for k := range p {
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pID == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func saveProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, products)
}

func updateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	product[pID] = reqBody.Name

	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		for k := range p {
			if pID == k {
				product = p
				index = i
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	products = splice(products, index)
	return c.JSON(http.StatusOK, product)
}

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "1323"
	}

	// Echo instance
	e := echo.New()
	e.Validator = &ProductValidator{validator: validator.New()}

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", home)

	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})
	e.GET("/products/:id", getProductById)
	e.POST("/products", saveProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	// Start server
	e.Logger.Print(fmt.Sprintf("access to http://localhost:%s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
