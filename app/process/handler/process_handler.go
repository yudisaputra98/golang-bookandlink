package handler

import (
	"github.com/labstack/echo/v4"
	services2 "github.com/yudisaputra/assignment-bookandlink/app/process/services"
)

type ProcessHandler struct{}

var processService = services2.NewProcessService()

func NewProcessHandler() *ProcessHandler {
	return &ProcessHandler{}
}

// all data
func (j *ProcessHandler) All(ctx echo.Context) error {
	data := processService.All()
	return ctx.JSON(data.Code, data)
}
