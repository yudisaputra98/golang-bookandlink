package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yudisaputra/assignment-bookandlink/app/job/handler"
)

var jobHandler = handler.NewJobHandler()

func Route() {
	r := echo.New()

	route := r.Group("api")
	{
		job := route.Group("/jobs")
		{
			job.GET("", jobHandler.All)
			job.GET("/:id", jobHandler.FindById)
			job.POST("/:id", jobHandler.Update)
			job.POST("", jobHandler.Create)
			job.GET("/:id", jobHandler.Delete)
		}
	}

	r.Logger.Fatal(r.Start(":8000"))
}
