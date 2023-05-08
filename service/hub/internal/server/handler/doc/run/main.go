package main

import (
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/doc"
)

func main() {
	config.Initialize()

	e := echo.New()
	e.GET("/", doc.New().OpenAPI())
	e.Logger.Fatal(e.Start(":3020"))
}
