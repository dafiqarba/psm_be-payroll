package services

import (
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type LeaveRecordService interface {
	//Read
	GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error)
	GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error)
	//Insert
	//InsertUser(user entity.User) (entity.User, error)
}

type leaveRecordService struct {
	leaveRecordRepository repository.LeaveRecordRepo
}

func NewLeaveRecordService(leaveRecordRepo repository.LeaveRecordRepo) LeaveRecordService {
	return &leaveRecordService{
		leaveRecordRepository: leaveRecordRepo,
	}
}

func (service *leaveRecordService) GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error) {
	return service.leaveRecordRepository.GetLeaveRecordDetail(req_id, id)
}

func (service *leaveRecordService) GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error) {
	if year == "ASC" || year == "DESC" {
		return service.leaveRecordRepository.GetLeaveRecordList(id, year)
	}
	year = "ASC"
	return service.leaveRecordRepository.GetLeaveRecordList(id, year)
}
