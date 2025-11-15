package models

import "gorm.io/gorm"

type StudentsTable struct {
	gorm.Model
	StudentID     string `json:"student_id" gorm:"unique; varchar(50); not null"`
	StudentPasswd string `json:"student_passwd" gorm:"unique; varchar(50); not null"`
	StudentName   string `json:"student_name" gorm:"varchar(30); not null"`
	Sex           int    `json:"sex" gorm:"type:tinyint not null"`
	Grade         string `json:"grade" gorm:"varchar(20)"`

	Classes []SelectableClasses `json:"classes" gorm:"many2many:Link; foreignKey:StudentID; joinForeignKey:StudentID; References:ClassID; joinReferences:ClassID; constraint:OnDelete:CASCADE"`
}
