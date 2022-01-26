package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type LeaveRecordRepo interface {
	//Read
	GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error)
	// GetLeaveRecordList() ([]entity.LeaveRecord, error)
	// GetLeaveBalance() ([]entity.LeaveBalance, error)
	//Create
	// InsertUser(user entity.User) (entity.User, error)
	//Update
	//Delete
}

type leaveRecordConnection struct {
	connection *sql.DB
}

func NewLeaveRecordRepo(dbConn *sql.DB) LeaveRecordRepo {
	return &leaveRecordConnection{
		connection: dbConn,
	}
}

func (db *leaveRecordConnection) GetLeaveRecordDetail(req_id int,id int) (entity.LeaveRecord, error) {
	// kita tutup koneksinya di akhir proses
	//defer db.connection.Close()
	//Variable to store collection of users
	var leaveRecordDetail entity.LeaveRecord
	//Execute SQL Query
	row := db.connection.QueryRow(`SELECT * FROM leave_records WHERE request_id=$1 AND user_id=$2`,req_id,id)
	err := row.Scan(&leaveRecordDetail.Request_id, &leaveRecordDetail.Request_on, &leaveRecordDetail.From_date, &leaveRecordDetail.To_date, &leaveRecordDetail.Return_date, &leaveRecordDetail.Reason, &leaveRecordDetail.Mobile, &leaveRecordDetail.Address, &leaveRecordDetail.Status_id, &leaveRecordDetail.Leave_id, &leaveRecordDetail.User_id)

	log.Println("|  Request received")
	//Error Handling
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return leaveRecordDetail, nil
	case nil:
		return leaveRecordDetail, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}	
	
	//Close the Execution of SQL Query
	//defer rows.Close()
	
	// return empty buku atau jika error
	return leaveRecordDetail, err
}
