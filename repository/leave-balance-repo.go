package repository

import (
	"database/sql"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type LeaveBalanceRepo interface {
	//Read
	GetLeaveBalance(id int, year string) (entity.LeaveBalance, error)
	// GetLeaveRecordList() ([]entity.LeaveRecord, error)
	// GetLeaveBalance() ([]entity.LeaveBalance, error)
	//Create
	// InsertUser(user entity.User) (entity.User, error)
	//Update
	//Delete
}

type leaveBalanceConnection struct {
	connection *sql.DB
}

func NewLeaveBalanceRepo(dbConn *sql.DB) LeaveBalanceRepo {
	return &leaveBalanceConnection{
		connection: dbConn,
	}
}

func (db *leaveBalanceConnection) GetLeaveBalance(id int, year string) (entity.LeaveBalance, error) {
	//Variable to store leave balance detail
	var leaveBalance entity.LeaveBalance
	//Execute SQL Query
	row := db.connection.QueryRow(`SELECT * FROM leave_balance WHERE user_id=$1 AND balance_year=$2`, id, year)
	err := row.Scan(&leaveBalance.Balance_id, &leaveBalance.Balance_year, &leaveBalance.Cuti_tahunan, &leaveBalance.Cuti_diambil, &leaveBalance.Cuti_balance, &leaveBalance.Cuti_izin, &leaveBalance.Cuti_sakit, &leaveBalance.User_id)

	//Err Handling
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("| "+err.Error())
			return leaveBalance, err
		} else {
			log.Println("| "+err.Error())
			return leaveBalance, err
		}
	}

	// returns populated data
	return leaveBalance, err
}
