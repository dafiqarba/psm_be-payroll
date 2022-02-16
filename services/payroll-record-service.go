package services

import (
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type PayrollRecordService interface {
	//Read
	GetPayrollRecordList() ([]entity.PayrollRecordListModel, error)
	GetPayrollRecordDetail(id int) (entity.PayrollRecordDetailModel, error)
	//Update
	// UpdatePayrollRecord(id int, p entity.PayrollRecord) (int, error)
	UpdatePayrollRecord(id int, p entity.PayrollRecord) (entity.PayrollRecord, error)
	//Create
	CreatePayrollRecord(p entity.PayrollRecord) (entity.PayrollRecord, error)
	// CreatePayrollRecord(p dto.CreatePayrollRecordModel) (int, error)
}

type payrollRecordService struct {
	payrollRecordRepo repository.PayrollRecordRepo
}

func NewPayrollRecordService(r repository.PayrollRecordRepo) PayrollRecordService {
	return &payrollRecordService{
		payrollRecordRepo: r,
	}
}

func (s *payrollRecordService) GetPayrollRecordList() ([]entity.PayrollRecordListModel, error) {
	return s.payrollRecordRepo.GetPayrollRecordList()
}

func (s *payrollRecordService) GetPayrollRecordDetail(id int) (entity.PayrollRecordDetailModel, error) {
	return s.payrollRecordRepo.GetPayrollRecordDetail(id)
}

func (s *payrollRecordService) CreatePayrollRecord(p entity.PayrollRecord) (entity.PayrollRecord, error) {
	// payrollRecord := entity.PayrollRecord{}
	// payrollRecord.Payment_period = p.Payment_period
	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
	// payrollRecord.Basic_salary = p.Basic_salary
	// payrollRecord.Bpjs = p.Bpjs
	// payrollRecord.Tax = p.Tax
	// payrollRecord.Total_salary = p.Total_salary
	// payrollRecord.Status_id = p.Status_id
	// payrollRecord.User_id = p.User_id

	return s.payrollRecordRepo.CreatePayrollRecord(p)
}

// func (s *payrollRecordService) CreatePayrollRecord(p dto.CreatePayrollRecordModel) (int, error) {
// 	var payrollRecord = entity.PayrollRecord{}
// 	payrollRecord.Payment_period = p.Payment_period
// 	payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
// 	payrollRecord.Basic_salary = p.Basic_salary
// 	payrollRecord.Bpjs = p.Bpjs
// 	payrollRecord.Tax = p.Tax
// 	payrollRecord.Total_salary = p.Total_salary
// 	payrollRecord.Status_id = p.Status_id
// 	payrollRecord.User_id = p.User_id

// 	return s.payrollRecordRepo.CreatePayrollRecord(payrollRecord)
// }

// func (s *payrollRecordService) UpdatePayrollRecord(p entity.PayrollRecord) (int, error) {
// 	var payrollRecord = entity.PayrollRecord{}
// 	payrollRecord.Payment_period = p.Payment_period
// 	payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
// 	payrollRecord.Basic_salary = p.Basic_salary
// 	payrollRecord.Bpjs = p.Bpjs
// 	payrollRecord.Tax = p.Tax
// 	payrollRecord.Total_salary = p.Total_salary
// 	payrollRecord.Status_id = p.Status_id
// 	payrollRecord.User_id = p.User_id

// 	return s.payrollRecordRepo.UpdatePayrollRecord(id, payrollRecord)
// }

func (s *payrollRecordService) UpdatePayrollRecord(id int, p entity.PayrollRecord) (entity.PayrollRecord, error) {
	// var payrollRecord = entity.PayrollRecord{}
	// payrollRecord.Payment_period = p.Payment_period
	// payrollRecord.Payment_date, _ = p.Payment_date, time.RFC822
	// payrollRecord.Basic_salary = p.Basic_salary
	// payrollRecord.Bpjs = p.Bpjs
	// payrollRecord.Tax = p.Tax
	// payrollRecord.Total_salary = p.Total_salary
	// payrollRecord.Status_id = p.Status_id
	// payrollRecord.User_id = p.User_id

	return s.payrollRecordRepo.UpdatePayrollRecord(id, p)
}
