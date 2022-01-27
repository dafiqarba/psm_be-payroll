package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
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
	
	var leaveBalance, err = c.leaveBalanceService.GetLeaveBalance(id, year)
	if err != nil {
		errMsg := errors.New(" the server cannot find the requested resource").Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg, http.StatusNotFound))
		return
		//response.Write([]byte(`{"error": Error getting the list"}`))
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(utils.ResponseJSON(http.StatusOK, "OK", leaveBalance))
}
