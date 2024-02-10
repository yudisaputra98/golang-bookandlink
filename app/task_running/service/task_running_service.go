package service

import (
	"github.com/yudisaputra/assignment-bookandlink/app/job/services"
	services2 "github.com/yudisaputra/assignment-bookandlink/app/process/services"
	"log"
	"runtime"
	"time"
)

var jobService = services.NewJobService()
var processService = services2.NewProcessService()

type TaskRunningServiceInterface interface {
	RunQueue()
}

type TaskRunningService struct{}

func NewTaskRunningService() TaskRunningServiceInterface {
	return &TaskRunningService{}
}

func (j *TaskRunningService) RunQueue() {
	start := time.Now()
	runtime.GOMAXPROCS(3)

	chanQueue := jobService.EnqueueJob()
	chanProcess := processService.ProcessJob(chanQueue)

	counterComplatePipeline := 0

	for chanP := range chanProcess {
		if chanP.Status == true {
			counterComplatePipeline++
		}
	}

	duration := time.Since(start)
	log.Println("done execute job in", duration.Seconds(), "seconds")
}
