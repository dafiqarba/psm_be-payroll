package repository

import (
	"database/sql"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type UserRepo interface {
	//Read
	GetUserList() ([]entity.User, error)
	GetUserDetail(id int) (entity.UserDetailModel, error)
	// GetLeaveRecordList() ([]entity.LeaveRecord, error)
	// GetLeaveBalance() ([]entity.LeaveBalance, error)
	//Create
	// InsertUser(user entity.User) (entity.User, error)
	//Update
	//Delete
}

type userConnection struct {
	connection *sql.DB
}

func NewUserRepo(dbConn *sql.DB) UserRepo {
	return &userConnection{
		connection: dbConn,
	}
}

func (db *userConnection) GetUserList() ([]entity.User, error) {
	//Variable to store collection of users
	var users []entity.User
	//Execute SQL Query
	rows, err := db.connection.Query(`SELECT * FROM users`)
	//Error Handling
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	log.Println("|  Requst received")
	//Close the Execution of SQL Query
	defer rows.Close()
	//Iterate over all available rows and strore the data
	for rows.Next() {
		var user entity.User
		// scan and assign into destination variable
		err = rows.Scan(&user.User_id, &user.Email, &user.Password, &user.Name, &user.Position_id, &user.Nik, &user.Role_id)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		// append to users slice
		users = append(users, user)
	}
	// returns populated data
	return users, err
}

func (db *userConnection) GetUserDetail(id int) (entity.UserDetailModel, error) {

	var userDetail entity.UserDetailModel
	//SQL Query
	query := `
		SELECT 
			u.user_id, u.name, u.position_id, u.nik, u.role_id, r.role_name, p.position_name 
		FROM users AS u 
			INNER JOIN roles AS r 
				ON r.role_id = u.role_id 
			INNER JOIN positions AS p 
				ON p.position_id = u.position_id 
		WHERE user_id=$1`
	//Execute SQL Query
	row := db.connection.QueryRow(query,id)
	err := row.Scan(&userDetail.User_id, &userDetail.Name, &userDetail.Position_id, &userDetail.Nik, &userDetail.Role_id, &userDetail.Role_name, &userDetail.Position_name)
	
	//Err Handling
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("| "+err.Error())
			return userDetail, err
		} else {
			log.Println("| "+err.Error())
			return userDetail, err
		}
	}
	
	// return empty buku atau jika error
	return userDetail, err
}