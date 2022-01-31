package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dafiqarba/be-payroll/dto"
	"github.com/dafiqarba/be-payroll/services"
	"github.com/dafiqarba/be-payroll/utils"
)

type AuthController interface {
	//Read Operation
	Login(response http.ResponseWriter, request *http.Request)
	//Create Operation
	Register(response http.ResponseWriter, request *http.Request)
}

type authController struct {
	authServ services.AuthService
	// jwtServ services.JWTService
	userServ services.UserService
}

func NewAuthController (authServ services.AuthService, userServ services.UserService) AuthController {
	return &authController {
		authServ : authServ,
		userServ: userServ,
	}
}

func (c *authController) Login(response http.ResponseWriter, request *http.Request) {
	var userLogin dto.UserLogin
	errDec := json.NewDecoder(request.Body).Decode(&userLogin)

	// Error handling
	if errDec != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errDec.Error())
		return
	}
	// Forwarding data to service
	var userLoginData, err = c.authServ.VerifyCredentials(userLogin)
	if err != nil {
		errMsg := errors.New("incorrect email/password").Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg, http.StatusUnauthorized))
		return
	}

		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(utils.ResponseJSON(http.StatusOK, "login success", userLoginData))

	
}

func (c *authController) Register(response http.ResponseWriter, request *http.Request) {
	//Var that holds registered user data
	var regUser dto.RegisterUser
	//Retrieve data from JSON
	errDec := json.NewDecoder(request.Body).Decode(&regUser)
	if errDec != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errDec.Error())
		return
	}
	//Forwarding data to user service
	createdUser, err := c.userServ.CreateUser(regUser)
	if err != nil {
		errMsg := err.Error()
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusConflict)
		json.NewEncoder(response).Encode(utils.ErrorJSON(errMsg, http.StatusConflict))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(utils.ResponseJSON(http.StatusOK,"OK",createdUser))
}