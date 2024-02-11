package main

import (
	"github.com/go-co-op/gocron"
	"github.com/labstack/gommon/log"
	"github.com/yudisaputra/assignment-bookandlink/app/task_running/service"
	"github.com/yudisaputra/assignment-bookandlink/core"
	"github.com/yudisaputra/assignment-bookandlink/database"
	"github.com/yudisaputra/assignment-bookandlink/routes"
	"time"
)

func main() {
	// load env
	core.LoadEnv()

	// load database
	database.Mysql()

	//load cron
	loc, err := time.LoadLocation("Asia/Makassar")
	taskRunningService := service.NewTaskRunningService()
	if err != nil {
		log.Error(err)
	}
	// start cronjob
	job := gocron.NewScheduler(loc)
	job.Every("5s").Do(taskRunningService.RunQueue)
	job.StartAsync()

	// route
	routes.Route()
}
