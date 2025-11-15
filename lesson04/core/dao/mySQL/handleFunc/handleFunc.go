package handleFunc

import (
	"errors"
	"lesson04/core"
	"lesson04/core/dao/mySQL/models"
	"lesson04/core/utils/md5"

	"gorm.io/gorm"
)

func (root *MySQLHandle) CheckStudentExists(stuID string) bool {
	tx := core.DB.Begin()
	err := tx.Where("student_id = ?", stuID).Find(&models.StudentsTable{}).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (root *MySQLHandle) AddStudent(stuData models.StudentsTable) error {

	data := models.StudentsTable{
		StudentID:     stuData.StudentID,
		StudentPasswd: md5.GenMD5(stuData.StudentPasswd),
		StudentName:   stuData.StudentName,
		Sex:           stuData.Sex,
		Grade:         stuData.Grade,
		Classes:       stuData.Classes,
	}

	tx := core.DB.Begin()
	err := tx.Create(&data).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func (root *MySQLHandle) UpdateStudent(stuData models.StudentsTable) error {
	tx := core.DB.Begin()
	err := tx.Model(&models.StudentsTable{}).Updates(&stuData).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func (root *MySQLHandle) DeleteStudent(stuID string) error {
	tx := core.DB.Begin()
	err := tx.Where("id = ?", stuID).Delete(&models.StudentsTable{}).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func (root *MySQLHandle) GetStudentInfo(stuID string) (models.StudentsTable, error) {
	stuForm := models.StudentsTable{}

	tx := core.DB.Begin()
	err := tx.Where("id = ?", stuID).Find(&stuForm).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return stuForm, err
}

func (root *MySQLHandle) GetSelectableClasses() ([]models.SelectableClasses, error) {
	classesForm := []models.SelectableClasses{}

	tx := core.DB.Begin()
	err := tx.Order("id asc").Find(&classesForm).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return classesForm, err
}

func (root *MySQLHandle) GetClassesOfStu(stuID string) (models.StudentsTable, error) {
	var student models.StudentsTable

	tx := core.DB.Begin()
	err := tx.Preload("Classes").Where("student_id = ?", stuID).Find(&student).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return student, err
}

func (root *MySQLHandle) SelectClassForStu(stuID string, classID string) error {
	var student models.StudentsTable
	var class models.SelectableClasses

	tx := core.DB.Begin()
	if err := tx.Where("student_id = ?", stuID).Find(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("class_id = ?", classID).Find(&class).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&student).Association("Classes").Append(&class); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (root *MySQLHandle) DeleteClassForStu(stuID string, classID string) error {
	var student models.StudentsTable
	var class models.SelectableClasses

	tx := core.DB.Begin()
	if err := tx.Where("student_id = ?", stuID).Find(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("class_id = ?", classID).Find(&class).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&student).Association("Classes").Delete(&class); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (root *MySQLHandle) CheckClassExists(stuID string, classID string) bool {
	tx := core.DB.Begin()
	err := tx.Where("student_id = ? AND class_id = ?", stuID, classID).Find(&models.Link{}).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return !errors.Is(err, gorm.ErrRecordNotFound)
}
