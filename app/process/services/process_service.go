package services

import (
	"github.com/labstack/gommon/log"
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	"github.com/yudisaputra/assignment-bookandlink/app/job/repository"
	"github.com/yudisaputra/assignment-bookandlink/app/job/services"
	entity2 "github.com/yudisaputra/assignment-bookandlink/app/process/entity"
	repository2 "github.com/yudisaputra/assignment-bookandlink/app/process/repository"
	"github.com/yudisaputra/assignment-bookandlink/helpers"
	"github.com/yudisaputra/assignment-bookandlink/responses"
	"gorm.io/gorm"
	"sync"
)

var processRepository = repository2.NewProcessRepository()
var jobRepository = repository.NewJobRepository()

type ProcessServiceInterface interface {
	All() responses.Api
	Create(data entity2.Process) responses.Api
	FindById(id string) responses.Api
	Update(id string, data entity2.Process) responses.Api
	Delete(id string) responses.Api
	ProcessJob(chanIn <-chan services.ChanResult) <-chan services.ChanResult
	MergeChan(chanInMany ...<-chan services.ChanResult) <-chan services.ChanResult
}

type ProcessService struct{}

func NewProcessService() ProcessServiceInterface {
	return &ProcessService{}
}

func (j *ProcessService) All() responses.Api {
	data, err := processRepository.All()

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: nil, Data: data}
}

func (j *ProcessService) Create(data entity2.Process) responses.Api {
	err := processRepository.Create(entity2.Process{
		ID:      helpers.Uid(16),
		JobName: data.JobName,
	})

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil disimpan"}
}

func (j *ProcessService) FindById(id string) responses.Api {
	data, err := processRepository.FindById(id)
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

func (j *ProcessService) Update(id string, data entity2.Process) responses.Api {
	err := processRepository.Update(id, entity2.Process{
		JobName: data.JobName,
	})

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil diubah"}
}

func (j *ProcessService) Delete(id string) responses.Api {
	err := processRepository.Delete(id)

	if err != nil {
		return responses.Api{Code: 400, Status: false, Message: nil, Error: err.Error(), Data: nil}
	}

	return responses.Api{Code: 200, Status: true, Message: "Job berhasil dihapus"}
}

func (j *ProcessService) MergeChan(chanInMany ...<-chan services.ChanResult) <-chan services.ChanResult {
	wg := new(sync.WaitGroup)
	chanOut := make(chan services.ChanResult)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan services.ChanResult) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func (j *ProcessService) ProcessJob(chanIn <-chan services.ChanResult) <-chan services.ChanResult {
	chanOut := make(chan services.ChanResult)

	go func() {
		jobs, err := processRepository.All()

		if err != nil {
			log.Error(err)
		}

		for chanProcess := range chanIn {
			if chanProcess.Status == true {
				for k, v := range jobs {
					// delete process
					err := processRepository.Delete(v.ID)

					if err != nil {
						log.Error(err)
					}

					// schema if fail process queue
					if k == 1 || k == 8 || k == 15 {
						job, _ := jobRepository.FindById(v.ID)

						err2 := jobRepository.Update(v.ID, entity.Job{
							Status:  "QUEUE",
							Attempt: job.Attempt + 1,
						})

						if err2 != nil {
							log.Error(err2)
						}

						err3 := processRepository.Create(entity2.Process{
							ID:      v.ID,
							JobName: v.JobName,
						})

						if err3 != nil {
							log.Error(err3)
						}
					} else {
						// schema if success
						// process queue
						err2 := jobRepository.Update(v.ID, entity.Job{
							Status: "SUCCESS",
						})

						if err2 != nil {
							log.Error(err2)
						}
					}

					//if 4 == k%5 {
					//	time.Sleep(2 * time.Second)
					//}
				}

				chanOut <- services.ChanResult{
					Status: true,
					Error:  nil,
				}
			}
		}

		close(chanOut)

	}()

	return chanOut
}
