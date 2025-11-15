package models

type Link struct {
	StudentID string `json:"student_id" gorm:"unique; varchar(50); not null"`
	ClassID   string `json:"class_id" gorm:"unique; not null; varchar(50)"`
}
