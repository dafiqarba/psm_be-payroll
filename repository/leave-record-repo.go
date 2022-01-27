package repository

import (
	"database/sql"
	"errors"
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
	//Variable to store leave record detail
	var leaveRecordDetail entity.LeaveRecord
	//Query
	query := `
		SELECT 
			* 
		FROM 
			leave_records 
		WHERE 
			request_id=$1 AND user_id=$2;
	`
	//Execute SQL Query
	row := db.connection.QueryRow(query, req_id, id)
	err := row.Scan(&leaveRecordDetail.Request_id, &leaveRecordDetail.Request_on, &leaveRecordDetail.From_date, &leaveRecordDetail.To_date, &leaveRecordDetail.Return_date, &leaveRecordDetail.Reason, &leaveRecordDetail.Mobile, &leaveRecordDetail.Address, &leaveRecordDetail.Status_id, &leaveRecordDetail.Leave_id, &leaveRecordDetail.User_id)

	//Err Handling
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("| "+err.Error())
			return leaveRecordDetail, err
		} else {
			log.Println("| "+err.Error())
			return leaveRecordDetail, err
		}
	}

	// returns populated data
	return leaveRecordDetail, err
}

func (db *leaveRecordConnection) GetLeaveRecordList(id int, year string) ([]entity.LeaveRecordListModel, error) {
	// ORDER BY request_on DESC
	var leaveRecordList []entity.LeaveRecordListModel
	//Execute SQL Query
	query := fmt.Sprintf(`
		SELECT 
			l.request_id, l.request_on, t.leave_name, l.reason, s.status_name, l.user_id 
		FROM 
			leave_records as l 
				INNER JOIN status as s 
					ON s.status_id = l.status_id 
				INNER JOIN leave_types as t 
					ON t.leave_id = l.leave_id 
		WHERE 
			l.user_id = %v ORDER BY l.request_on %v;`, id, year)

	rows, err := db.connection.Query(query)
	if err != nil {
		log.Fatalf("cannot execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var leaveRecord entity.LeaveRecordListModel
		// scan and assign into leaveRecord
		err = rows.Scan(&leaveRecord.Request_id, &leaveRecord.Request_on, &leaveRecord.Leave_name, &leaveRecord.Reason, &leaveRecord.Status_name, &leaveRecord.User_id)
		if err != nil {
			log.Fatalf("cannot retrieve the data. %v", err)
		}
		// append to leaveRecordList slices
		leaveRecordList = append(leaveRecordList, leaveRecord)
	}
	// Check for empty result
	if len(leaveRecordList) == 0 {
		log.Println("| "+errors.New("sql: no results").Error())
		err := errors.New("sql: no results")
		return leaveRecordList, err
	}
	// return leaveRecordlist populated with results
	return leaveRecordList, err
}
