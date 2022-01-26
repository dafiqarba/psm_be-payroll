package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dafiqarba/be-payroll/config"
	"github.com/dafiqarba/be-payroll/controller"
	"github.com/dafiqarba/be-payroll/repository"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/gorilla/mux"
)

var (
	db             *sql.DB                   = config.SetupDatabaseConnection()

	userRepository repository.UserRepo       = repository.NewUserRepo(db)
	userService    services.UserService      = services.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)

	leaveBalanceRepository repository.LeaveBalanceRepo = repository.NewLeaveBalanceRepo(db)
	leaveBalanceService services.LeaveBalanceService = services.NewLeaveBalanceService(leaveBalanceRepository)
	leaveBalanceController controller.LeaveBalanceController = controller.NewLeaveBalanceController(leaveBalanceService)

	leaveRecordRepository repository.LeaveRecordRepo = repository.NewLeaveRecordRepo(db)
	leaveRecordService services.LeaveRecordService = services.NewLeaveRecordService(leaveRecordRepository)
	leaveRecordController controller.LeaveRecordController = controller.NewLeaveRecordController(leaveRecordService)
)

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/user-detail", getAll).Methods(http.MethodGet)
	router.HandleFunc("/user-list", userController.GetUserList).Methods(http.MethodGet)
	//URL Pattern: localhost:8000/leave-balance?id=1&year=2022
	router.HandleFunc("/leave-balance", leaveBalanceController.GetLeaveBalance).Methods(http.MethodGet)
	//URL Pattern: localhost:8000/leave-record-detail?req_id=1&id=2
	router.HandleFunc("/leave-record-detail", leaveRecordController.GetLeaveRecordDetail).Methods(http.MethodGet)

	//Start server
	log.Println("|  Server listening on port: 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))

	// VERSION 1 HARDCODED WITH NET/HTTP PACKAGE
	// mux := http.NewServeMux()
	// //Routes definition
	// mux.HandleFunc("/user-detail", getAll)
	// //Start server
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

// type User struct {
// 	Name string `json:"name"`
// 	City string `json:"city"`
// }

// func getAll(w http.ResponseWriter, r *http.Request) {

// 	users := []User{
// 		{Name: "Diane", City: "LA"},
// 		{Name: "Podolski", City: "Koln"},
// 	}
// 	js, err := json.MarshalIndent(users, "", "\t")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(js)
// 	// json.NewEncoder(w).Encode(users)
// }

// GO WITH REACT:

// http.HandleFunc("/status", func( w http.ResponseWriter, r *http.Request) {
// 	currentStatus := config.AppStatus {
// 		Status: "Available",
// 		Environment: "Development",
// 		Version: config.Version,
// 	}
// 	js, err := json.MarshalIndent(currentStatus, "", "\t")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// })

// err := http.ListenAndServe()
