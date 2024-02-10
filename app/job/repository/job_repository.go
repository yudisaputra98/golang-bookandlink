package repository

import (
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	"github.com/yudisaputra/assignment-bookandlink/database"
)

type JobRepositoryInterface interface {
	All() (entity.Jobs, error)
	Create(data entity.Job) error
	FindById(id string) (entity.Job, error)
	Update(id string, data entity.Job) error
	Delete(id string) error
	GetNotDone() (entity.Jobs, error)
}

type JobRepository struct{}

func NewJobRepository() JobRepositoryInterface {
	return &JobRepository{}
}

func (j *JobRepository) All() (entity.Jobs, error) {
	var jobs entity.Jobs

	err := database.Instance.Order("created_at desc").Find(&jobs)
	return jobs, err.Error
}

func (j *JobRepository) Create(data entity.Job) error {
	err := database.Instance.Create(&data)
	return err.Error
}

func (j *JobRepository) FindById(id string) (entity.Job, error) {
	var job entity.Job

	err := database.Instance.Where("id = ?", id).First(&job)
	return job, err.Error
}

func (j *JobRepository) Update(id string, data entity.Job) error {
	err := database.Instance.Where("id = ?", id).Updates(&data)
	return err.Error
}

func (j *JobRepository) Delete(id string) error {
	var job entity.Job

	err := database.Instance.Delete(&job, "id = ?", id)
	return err.Error
}

func (j *JobRepository) GetNotDone() (entity.Jobs, error) {
	var jobs entity.Jobs

	err := database.Instance.Where("status = 0").Order("created_at desc").Find(&jobs)
	return jobs, err.Error
}
