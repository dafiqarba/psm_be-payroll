package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
)

type UserController interface {
	//Read Operation
	GetUserList(response http.ResponseWriter, request *http.Request)
	GetUserDetail(response http.ResponseWriter, request *http.Request)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userServ services.UserService) UserController {
	return &userController{
		userService: userServ,
	}
}

func (c *userController) GetUserList(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var users, err = c.userService.GetUserList()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(`{"error": "Error getting the list"}`)
		//response.Write([]byte(`{"error": Error getting the list"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func (c *userController) GetUserDetail(response http.ResponseWriter, request *http.Request) {
	v := request.URL.Query()
	id,_ := strconv.Atoi(v.Get("id"))

	var userDetail, err = c.userService.GetUserDetail(id)
	if err != nil {
		errMsg := errors.New("the server cannot find the requested resource").Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg, http.StatusNotFound))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(utils.ResponseJSON(http.StatusOK,"OK",userDetail))
}
