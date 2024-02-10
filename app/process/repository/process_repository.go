package repository

import (
	entity2 "github.com/yudisaputra/assignment-bookandlink/app/process/entity"
	"github.com/yudisaputra/assignment-bookandlink/database"
)

type ProcessRepositoryInterface interface {
	All() (entity2.Processes, error)
	Create(data entity2.Process) error
	FindById(id string) (entity2.Process, error)
	Update(id string, data entity2.Process) error
	Delete(id string) error
}

type ProcessRepository struct{}

func NewProcessRepository() ProcessRepositoryInterface {
	return &ProcessRepository{}
}

func (j *ProcessRepository) All() (entity2.Processes, error) {
	var processes entity2.Processes

	err := database.Instance.Order("created_at desc").Find(&processes)
	return processes, err.Error
}

func (j *ProcessRepository) Create(data entity2.Process) error {
	err := database.Instance.Create(&data)
	return err.Error
}

func (j *ProcessRepository) FindById(id string) (entity2.Process, error) {
	var process entity2.Process

	err := database.Instance.Where("id = ?", id).First(&process)
	return process, err.Error
}

func (j *ProcessRepository) Update(id string, data entity2.Process) error {
	err := database.Instance.Where("id = ?", id).Updates(&data)
	return err.Error
}

func (j *ProcessRepository) Delete(id string) error {
	var process entity2.Process

	err := database.Instance.Delete(&process, "id = ?", id)
	return err.Error
}
