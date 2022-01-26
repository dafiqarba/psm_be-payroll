package repository

import (
	"database/sql"
	"fmt"
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
	// kita tutup koneksinya di akhir proses
	//defer db.connection.Close()
	//Variable to store collection of users
	var leaveBalance entity.LeaveBalance
	//Execute SQL Query
	row := db.connection.QueryRow(`SELECT * FROM leave_balance WHERE user_id=$1 AND balance_year=$2`,id,year)
	err := row.Scan(&leaveBalance.Balance_id, &leaveBalance.Balance_year, &leaveBalance.Cuti_tahunan, &leaveBalance.Cuti_diambil, &leaveBalance.Cuti_balance, &leaveBalance.Cuti_izin, &leaveBalance.Cuti_sakit, &leaveBalance.User_id)

	//Error Handling
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return leaveBalance, nil
	case nil:
		return leaveBalance, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	log.Println("|  Request received")
	//Close the Execution of SQL Query
	//defer rows.Close()
	
	// return empty buku atau jika error
	return leaveBalance, err
}
