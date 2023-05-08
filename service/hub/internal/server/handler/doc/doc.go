package doc

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/pkg/jschema"
)

type Doc struct {
	schemas jschema.Schemas
}

type Obj map[string]interface{}

func New() *Doc {
	return &Doc{
		schemas: jschema.New("#/components/schemas"),
	}
}

func (d *Doc) Generate() interface{} {
	d.Define()

	return Obj{
		"openapi": "3.0.1",
		"info": Obj{
			"title":   "Pregod API",
			"version": "current",
		},
		"paths": d.endpoints(),
		"components": Obj{
			"schemas": d.schemas.JSON(),
		},
	}
}

func (d *Doc) OpenAPI() echo.HandlerFunc {
	cache := d.Generate()
	return func(c echo.Context) error {
		if c.Request().URL.Query().Get("json") != "" {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			return c.JSON(http.StatusOK, cache)
		}

		proto := "http://"
		v := ""

		if strings.Contains(c.Request().Host, "pregod.rss3.dev") {
			proto = "https://"
			v = "/v1"
		}

		u := url.QueryEscape(proto + c.Request().Host + v + c.Request().URL.Path + "?json=true")
		fmt.Println("url", u)

		return c.Redirect(http.StatusFound, fmt.Sprintf("https://generator.swagger.io/?url=%s", u))
	}
}
