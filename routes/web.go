package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yudisaputra/assignment-bookandlink/app/job/handler"
	handler2 "github.com/yudisaputra/assignment-bookandlink/app/process/handler"
)

var jobHandler = handler.NewJobHandler()
var ProcessHandler = handler2.NewProcessHandler()

func Route() {
	r := echo.New()

	route := r.Group("api")
	{
		job := route.Group("/jobs")
		{
			job.GET("", jobHandler.All)
			job.POST("", jobHandler.Create)
			job.GET("/generate", jobHandler.Generate)
			job.GET("/:id", jobHandler.FindById)
			job.POST("/:id", jobHandler.Update)
			job.GET("/:id", jobHandler.Delete)
		}
		process := route.Group("/processes")
		{
			process.GET("", ProcessHandler.All)
		}
	}

	r.Logger.Fatal(r.Start(":8000"))
}
