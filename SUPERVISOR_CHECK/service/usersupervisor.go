package service

import (
	"SUPERVISOR_CHECK/model"
	"SUPERVISOR_CHECK/repository"
)

type UserSupervisorsService interface {
	GetSupervisors(Department string) ([]model.UserSupervisors, error)
}

type userSupervisorService struct {
	ur repository.UserSupervisorsRepository
}

func NewUserSupervisor(ur repository.UserSupervisorsRepository) UserSupervisorsService {
	return &userSupervisorService{ur}
}

func (us *userSupervisorService) GetSupervisors(Department string) ([]model.UserSupervisors, error) {
	return us.ur.GetSupervisors(Department)
}
