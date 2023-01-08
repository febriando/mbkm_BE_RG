package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *UserRepository) UserAvail(user model.User) error {
	// checkField := `select * from users` db.Model(&User{}).Take(&result)
	// result := user{}
	// if err := u.db.Take(&user).Error; err != nil {
	// 	return err
	// }
	if err := u.db.
		Where("username = ? AND password = ?", user.Username, user.Password).
		First(&user).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *UserRepository) CheckPassLength(pass string) bool {
	if len(pass) <= 5 {
		return true
	}

	return false
}

func (u *UserRepository) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
