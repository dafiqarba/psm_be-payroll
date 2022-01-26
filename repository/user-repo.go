package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type UserRepo interface {
	//Read
	GetUserList() ([]entity.User, error)
	GetUserDetail(id int) (entity.User, error)
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
	// kita tutup koneksinya di akhir proses
	//defer db.connection.Close()
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
	//defer rows.Close()
	//Iterate over all available rows and strore the data
	for rows.Next() {
		var user entity.User
		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&user.User_id, &user.Email, &user.Password, &user.Name, &user.Position_id, &user.Nik, &user.Role_id)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		// masukkan kedalam slice bukus
		users = append(users, user)
	}
	// return empty buku atau jika error
	return users, err
}

func (db *userConnection) GetUserDetail(id int) (entity.User, error) {

	var userDetail entity.User
	//Execute SQL Query
	row := db.connection.QueryRow(`SELECT user_id, name, position_id, nik, role_id FROM users WHERE user_id=$1`,id)
	err := row.Scan(&userDetail.User_id, &userDetail.Name, &userDetail.Position_id, &userDetail.Nik, &userDetail.Role_id)
	
	log.Println("|  Request received")
	//Error Handling
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return userDetail, nil
	case nil:
		return userDetail, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}	
	
	// return empty buku atau jika error
	return userDetail, err
}