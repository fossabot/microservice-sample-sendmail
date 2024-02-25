package repository

import (
	"SUPERVISOR_CHECK/model"

	"gorm.io/gorm"
)

type UserSupervisorsRepository interface {
	GetSupervisors(Department string) ([]model.UserSupervisors, error)
}

type userSupervisorsRepository struct {
	db *gorm.DB
}

func NewUserSupervisor(db *gorm.DB) UserSupervisorsRepository {
	return &userSupervisorsRepository{db}
}

func (u *userSupervisorsRepository) GetSupervisors(Department string) ([]model.UserSupervisors, error) {
	var supervisors []model.UserSupervisors

	result := u.db.Table("users").
		Select("users.user_id, users.lastname, users.firstname, users.email, users.position_no, users.department_no, departments.department, departments.section, positions.position_name, positions.rankvalue").
		Joins("INNER JOIN departments ON users.department_no = departments.department_no").
		Joins("INNER JOIN positions ON users.position_no = positions.position_no").
		Where("departments.department = ?", Department).
		Where("positions.rankvalue > 0").
		Find(&supervisors)

	if result.Error != nil {
		return nil, result.Error
	}

	return supervisors, nil
}
