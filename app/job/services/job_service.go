package services

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	"github.com/yudisaputra/assignment-bookandlink/app/job/repository"
	entity2 "github.com/yudisaputra/assignment-bookandlink/app/process/entity"
	repository2 "github.com/yudisaputra/assignment-bookandlink/app/process/repository"
	"github.com/yudisaputra/assignment-bookandlink/helpers"
	"github.com/yudisaputra/assignment-bookandlink/responses"
	"gorm.io/gorm"
	"time"
)

var jobRepository = repository.NewJobRepository()
var processRepository = repository2.NewProcessRepository()

type ChanResult struct {
	Status bool
	Error  error
}

type JobServiceInterface interface {
	All(status int) responses.Api
	Create(data entity.Job) responses.Api
	FindById(id string) responses.Api
	Update(id string, data entity.Job) responses.Api
	Delete(id string) responses.Api
	Generate(total int) responses.Api
	EnqueueJob() <-chan ChanResult
}

type JobService struct{}

func NewJobService() JobServiceInterface {
	return &JobService{}
}

func (j *JobService) All(status int) responses.Api {
	data, err := jobRepository.All(status)

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: nil, Data: data}
}

func (j *JobService) Create(data entity.Job) responses.Api {
	err := jobRepository.Create(entity.Job{
		ID:      helpers.Uid(16),
		JobName: data.JobName,
		Status:  data.Status,
	})

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil disimpan"}
}

func (j *JobService) FindById(id string) responses.Api {
	data, err := jobRepository.FindById(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return responses.Api{Code: 404, Status: false, Message: "Data tidak ditemukan", Error: err.Error(), Data: nil}
		default:
			return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
		}
	}

	return responses.Api{Code: 200, Status: true, Data: data}
}

func (j *JobService) Update(id string, data entity.Job) responses.Api {
	err := jobRepository.Update(id, entity.Job{
		JobName: data.JobName,
		Status:  data.Status,
	})

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil diubah"}
}

func (j *JobService) Delete(id string) responses.Api {
	err := jobRepository.Delete(id)

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil dihapus"}
}

func (j *JobService) Generate(total int) responses.Api {
	//database.Instance.Migrator().DropTable("jobs")
	//database.Instance.AutoMigrate(&entity.Job{})

	for i := 1; i <= total; i++ {
		err := jobRepository.Create(entity.Job{
			ID:      helpers.Uid(16),
			JobName: fmt.Sprint("Generate job ", helpers.Uid(20)),
			Status:  0,
		})

		if err != nil {
			return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
		}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil dibuat"}
}

func (j *JobService) EnqueueJob() <-chan ChanResult {
	chanOut := make(chan ChanResult)

	go func() {
		jobs, err := jobRepository.GetNotDone()

		if err != nil {
			log.Error(err)
		}

		for k, v := range jobs {
			err := processRepository.Create(entity2.Process{
				ID:      v.ID,
				JobName: v.JobName,
			})

			if err != nil {
				log.Error(err)
			}

			err2 := jobRepository.Update(v.ID, entity.Job{
				Status: 1,
			})

			if err2 != nil {
				log.Error(err2)
			}

			if 4 == k%5 {
				time.Sleep(2 * time.Second)
			}
		}

		chanOut <- ChanResult{
			Status: true,
			Error:  nil,
		}

		close(chanOut)

	}()

	return chanOut
}
