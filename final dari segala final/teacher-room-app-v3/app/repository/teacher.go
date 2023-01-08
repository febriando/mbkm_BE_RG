package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	if err := u.db.Create(&teacher); err != nil {
		return err.Error
	}

	return nil // TODO: replace this
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	teacher := []model.ViewTeacher{}

	if err := u.db.Table("teachers").Where("deleted_at IS NULL").Find(&teacher); err != nil {
		return nil, err.Error
	}

	return nil, nil // TODO: replace this
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	if err := u.db.Table("teachers").Where("id = ? AND name = ?", id, name).Update("id", id).Error; err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	if err := u.db.Table("teachers").Where("id = ?", id).Delete("id", id).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
