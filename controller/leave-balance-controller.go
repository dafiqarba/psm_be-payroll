package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/services"
)

type LeaveBalanceController interface {
	GetLeaveBalance(response http.ResponseWriter, request *http.Request)
}

type leaveBalanceController struct {
	leaveBalanceService services.LeaveBalanceService
}

func NewLeaveBalanceController(leaveBalanceServ services.LeaveBalanceService) LeaveBalanceController {
	return &leaveBalanceController{
		leaveBalanceService: leaveBalanceServ,
	}
}

func (c *leaveBalanceController) GetLeaveBalance(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	id,_ := strconv.Atoi(v.Get("id"))
	year := v.Get("year")
	response.Header().Set("Content-Type", "application/json")
	var leaveBalance, err = c.leaveBalanceService.GetLeaveBalance(id, year)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(`{"error": "Error getting the data"}`)
		//response.Write([]byte(`{"error": Error getting the list"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(leaveBalance)
}
