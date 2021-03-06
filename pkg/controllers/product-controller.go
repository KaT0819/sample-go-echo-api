package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

var products = []map[int]string{{1: "mobilies"}, {2: "TV"}, {3: "laptop"}}

// Validator
func SetValidator(e *echo.Echo) {
	e.Validator = &ProductValidator{validator: validator.New()}
}

func (p *ProductValidator) Validate(i interface{}) error {
	if err := p.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Handler
func Home(c echo.Context) error {
	return c.HTML(http.StatusOK, `
<html lang="ja">
<head><title>Document</title></head>
<body> Hello Echo </body>
</html>
	`)
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func GetProductById(c echo.Context) error {
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

func SaveProduct(c echo.Context) error {
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

func UpdateProduct(c echo.Context) error {
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

func DeleteProduct(c echo.Context) error {
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
