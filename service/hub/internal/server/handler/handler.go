package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	utils "github.com/naturalselectionlabs/pregod/common/utils/interface"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/service"
)

type Handler struct {
	service *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) apiReport(path string, c echo.Context) {
	report := map[string]interface{}{
		"index":       model.EsIndex,
		"path":        path,
		"ts":          time.Now().Format("2006-01-02 15:04:05"),
		"count":       true,
		"api_key":     c.Get("API-KEY"),
		"remote_addr": c.RealIP(),
	}

	output, _ := json.Marshal(report)
	fmt.Printf("[DATABEAT]%v\n", string(output))
}

func (h *Handler) filterReport(path string, request interface{}) {
	b, _ := json.Marshal(request)
	report := make(map[string]interface{})

	err := json.Unmarshal(b, &report)
	if err != nil {
		return
	}

	for k, v := range report {
		if utils.IfInterfaceValueIsNil(v) {
			delete(report, k)
		}
	}

	if path == model.PostNotes {
		report["index"] = model.EsIndex
		report["path"] = path
		report["ts"] = time.Now().Format("2006-01-02 15:04:05")
		for _, address := range report["address"].([]interface{}) {
			report["address"] = address

			output, _ := json.Marshal(report)
			fmt.Printf("[DATABEAT]%v\n", string(output))
		}
	}
}
