package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
	_ "github.com/lib/pq"
)

type LeaveRecordRepo interface {
	//Read
	GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error)
	GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error)
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

func (db *leaveRecordConnection) GetLeaveRecordDetail(req_id int, id int) (entity.LeaveRecord, error) {
	// kita tutup koneksinya di akhir proses
	//defer db.connection.Close()
	//Variable to store collection of users
	var leaveRecordDetail entity.LeaveRecord
	//Execute SQL Query
	row := db.connection.QueryRow(`SELECT * FROM leave_records WHERE request_id=$1 AND user_id=$2`, req_id, id)
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

func (db *leaveRecordConnection) GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error) {
	// ORDER BY request_on DESC
	var leaveRecordList []entity.LeaveRecordListModel
	//Execute SQL Query
	query := fmt.Sprintf("SELECT l.request_id, l.request_on, t.leave_name, l.reason, s.status_name, l.user_id FROM leave_records as l INNER JOIN status as s ON s.status_id = l.status_id INNER JOIN leave_types as t ON t.leave_id = l.leave_id WHERE l.user_id = %v ORDER BY l.request_on %v;", id, year)

	rows, err := db.connection.Query(query)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	log.Println("|  Requst received")

	for rows.Next() {
		var leaveRecord entity.LeaveRecordListModel
		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&leaveRecord.Request_id, &leaveRecord.Request_on, &leaveRecord.Leave_name, &leaveRecord.Reason, &leaveRecord.Status_name, &leaveRecord.User_id)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		// masukkan kedalam slice bukus
		leaveRecordList = append(leaveRecordList, leaveRecord)
	}
	// return empty leaveRecordlist atau jika error
	return leaveRecordList, err
}
