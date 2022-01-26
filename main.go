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
	db *sql.DB = config.SetupDatabaseConnection()

	userRepository repository.UserRepo       = repository.NewUserRepo(db)
	userService    services.UserService      = services.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)

	leaveBalanceRepository repository.LeaveBalanceRepo       = repository.NewLeaveBalanceRepo(db)
	leaveBalanceService    services.LeaveBalanceService      = services.NewLeaveBalanceService(leaveBalanceRepository)
	leaveBalanceController controller.LeaveBalanceController = controller.NewLeaveBalanceController(leaveBalanceService)

	leaveRecordRepository repository.LeaveRecordRepo       = repository.NewLeaveRecordRepo(db)
	leaveRecordService    services.LeaveRecordService      = services.NewLeaveRecordService(leaveRecordRepository)
	leaveRecordController controller.LeaveRecordController = controller.NewLeaveRecordController(leaveRecordService)
)

func main() {
	router := mux.NewRouter()

	/*-----------------------------------------------------------------------
		User Route:

		User Detail				= /user-detail?id=2
		Leave Balance Detail	= /leave-balance?id=1&year=2022
		Leave Record Detail		= /leave-record-detail?req_id=1&id=2
		Leave Record List		= /leave-record-list?id=2&year=ASC
		Create Leave Record		= ?
	 ------------------------------------------------------------------------*/
	router.HandleFunc("/user-list", userController.GetUserList).Methods(http.MethodGet)
	router.HandleFunc("/leave-balance", leaveBalanceController.GetLeaveBalance).Methods(http.MethodGet)
	router.HandleFunc("/leave-record-detail", leaveRecordController.GetLeaveRecordDetail).Methods(http.MethodGet)
	router.HandleFunc("/leave-record-list", leaveRecordController.GetLeaveRecordList).Methods(http.MethodGet)
	router.HandleFunc("/user-detail", userController.GetUserDetail).Methods(http.MethodGet)

	//Start server
	log.Println("|  Server listening on port: 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
