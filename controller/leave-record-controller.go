package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
)

type LeaveRecordController interface {
	GetLeaveRecordDetail(response http.ResponseWriter, request *http.Request)
	GetLeaveRecordList(response http.ResponseWriter, request *http.Request)
	CreateLeaveRecord(response http.ResponseWriter, request *http.Request)
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

	
	var leaveRecordDetail, err = c.leaveRecordService.GetLeaveRecordDetail(req_id, id)

	if err != nil {
		errMsg := errors.New(" the server cannot find the requested resource").Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg, http.StatusNotFound))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(utils.ResponseJSON(http.StatusOK,"OK", leaveRecordDetail))
}

func (c *leaveRecordController) GetLeaveRecordList(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	id,_ := strconv.Atoi(v.Get("id"))
	year := v.Get("year")

	var leaveRecordList, err = c.leaveRecordService.GetLeaveRecordList(id, year)
	if err != nil {
		errMsg := errors.New("the server cannot find the requested resource").Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg,http.StatusNotFound))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(utils.ResponseJSON(http.StatusOK,"OK",leaveRecordList))
}

func (c *leaveRecordController) CreateLeaveRecord(response http.ResponseWriter, request *http.Request) {
	// Reference to CreateLeaveRecord data transfer obj
	var createLeaveRecord dto.CreateLeaveRecordModel
	// Retrieve body obj from request
	errDec := json.NewDecoder(request.Body).Decode(&createLeaveRecord)
	// Error handling
	if errDec != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errDec.Error())
		return
	}
	// Forwarding data to service
	var req_id, err = c.leaveRecordService.CreateLeaveRecord(createLeaveRecord)
	if err != nil {
		errMsg := errors.New("internal Server Error").Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg,http.StatusInternalServerError))
		return 
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(utils.InsertResponseJSON(http.StatusCreated,"Request created",req_id))
}
