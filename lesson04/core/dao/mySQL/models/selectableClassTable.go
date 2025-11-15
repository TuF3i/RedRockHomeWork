package models

import (
	"gorm.io/gorm"
)

type SelectableClasses struct {
	gorm.Model
	ClassName     string `json:"class_name" gorm:"not null; varchar(30)"`
	ClassID       string `json:"class_id" gorm:"unique; not null; varchar(50)"`
	ClassTeacher  string `json:"class_teacher" gorm:"not null; varchar(30)"`
	ClassTime     string `json:"class_time" gorm:"not null; varchar(50)"`
	SelectableNum int    `json:"selectable_num" gorm:"not null; type:tinyint"`
	SelectedNum   int    `json:"selected_num" gorm:"not null; type:tinyint"`

	Students []StudentsTable `json:"students" gorm:"many2many:Link; foreignKey:ClassID; joinForeignKey:ClassID; References:StudentID; joinReferences:StudentID"`
}
