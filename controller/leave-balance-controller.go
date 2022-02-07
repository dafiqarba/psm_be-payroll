package controller

import (
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
		utils.BuildErrorResponse(response, http.StatusNotFound, errMsg)
		return
	}
	utils.BuildResponse(response, http.StatusOK, "success", leaveBalance)
}
