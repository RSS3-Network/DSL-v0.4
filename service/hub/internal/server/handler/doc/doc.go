package doc

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/pkg/jschema"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
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

		proto := "https://"
		if config.ConfigHub.Mode == "development" {
			proto = "http://"
		}

		u := url.QueryEscape(proto + c.Request().Host + c.Request().URL.Path + "?json=true")
		fmt.Println("url", u)

		return c.Redirect(http.StatusFound, fmt.Sprintf("https://generator.swagger.io//?url=%s", u))
	}
}
