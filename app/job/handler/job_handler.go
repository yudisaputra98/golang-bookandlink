package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	"github.com/yudisaputra/assignment-bookandlink/app/job/services"
	services2 "github.com/yudisaputra/assignment-bookandlink/app/process/services"
	"github.com/yudisaputra/assignment-bookandlink/responses"
	"net/http"
)

type JobHandler struct{}

var jobService = services.NewJobService()
var processService = services2.NewProcessService()

func NewJobHandler() *JobHandler {
	return &JobHandler{}
}

// all data
func (j *JobHandler) All(ctx echo.Context) error {
	data := jobService.All()
	return ctx.JSON(data.Code, data)
}

// create data
func (j *JobHandler) Create(ctx echo.Context) error {
	var job entity.Job

	err := ctx.Bind(&job)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Api{Status: false, Code: 400, Message: nil, Error: err, Data: nil})
	}

	data := jobService.Create(job)
	return ctx.JSON(data.Code, data)
}

// find data
func (j *JobHandler) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	data := jobService.FindById(id)

	return ctx.JSON(data.Code, data)
}

// update data
func (j *JobHandler) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	var job entity.Job

	err := ctx.Bind(&job)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Api{Status: false, Code: 400, Message: nil, Error: err, Data: nil})
	}

	data := jobService.Update(id, job)
	return ctx.JSON(data.Code, data)
}

// delete data
func (j *JobHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	data := jobService.Delete(id)

	return ctx.JSON(data.Code, data)
}

func (j *JobHandler) Generate(ctx echo.Context) error {
	data := jobService.Generate(50)
	return ctx.JSON(data.Code, data)
}
