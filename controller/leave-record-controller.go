package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/services"
)

type LeaveRecordController interface {
	GetLeaveRecordDetail(response http.ResponseWriter, request *http.Request)
}

type leaveRecordController struct {
	leaveRecordService services.LeaveRecordService
}

func NewLeaveRecordController(leaveRecordServ services.LeaveRecordService) LeaveRecordController {
	return &leaveRecordController{
		leaveRecordService: leaveRecordServ,
	}
}

func (c *leaveRecordController) GetLeaveRecordDetail(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	req_id,_ := strconv.Atoi(v.Get("req_id"))
	id,_ := strconv.Atoi(v.Get("id"))
	response.Header().Set("Content-Type", "application/json")
	var leaveBalance, err = c.leaveRecordService.GetLeaveRecordDetail(req_id, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(`{"error": "Error getting the data"}`)
		//response.Write([]byte(`{"error": Error getting the list"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(leaveBalance)
}
