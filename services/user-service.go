package services

import (
	"github.com/dafiqarba/be-payroll/entity"
	"github.com/dafiqarba/be-payroll/repository"
)

type UserService interface {
	//Read
	GetUserList() ([]entity.User, error)
	//Insert
	//InsertUser(user entity.User) (entity.User, error)
}

type userService struct{
	userRepository repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService {
		userRepository: userRepo,
	}
}

func (service *userService) GetUserList() ([]entity.User, error) {
	return service.userRepository.GetUserList()
}
 


// From REST-BASED MICROSERVICE
// type DefaultUserService struct {
// 	repo repository.UserRepo
// }

// func (s DefaultUserService) GetUserList() ([]entity.User, error) {
// 	return s.repo.GetUserList()
// }

// func NewUserService(repository repository.UserRepo) DefaultUserService {
// 	return DefaultUserService{repository}
// }