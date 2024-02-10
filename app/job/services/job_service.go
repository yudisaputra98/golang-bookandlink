package services

import (
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	"github.com/yudisaputra/assignment-bookandlink/app/job/repository"
	"github.com/yudisaputra/assignment-bookandlink/helpers"
	"github.com/yudisaputra/assignment-bookandlink/responses"
	"gorm.io/gorm"
)

var jobRepository = repository.NewJobRepository()

type JobServiceInterface interface {
	All() responses.Api
	Create(data entity.Job) responses.Api
	FindById(id string) responses.Api
	Update(id string, data entity.Job) responses.Api
	Delete(id string) responses.Api
}

type JobService struct{}

func NewJobService() JobServiceInterface {
	return &JobService{}
}

func (j *JobService) All() responses.Api {
	data, err := jobRepository.All()

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
