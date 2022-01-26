package services

import (
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type LeaveBalanceService interface {
	//Read
	GetLeaveBalance(id int, year string) (entity.LeaveBalance, error)
	//Insert
	//InsertUser(user entity.User) (entity.User, error)
}

type leaveBalanceService struct {
	leaveBalanceRepository repository.LeaveBalanceRepo
}

func NewLeaveBalanceService(leaveBalanceRepo repository.LeaveBalanceRepo) LeaveBalanceService {
	return &leaveBalanceService{
		leaveBalanceRepository: leaveBalanceRepo,
	}
}

func (service *leaveBalanceService) GetLeaveBalance(id int, year string) (entity.LeaveBalance,error) {
	return service.leaveBalanceRepository.GetLeaveBalance(id, year)
}