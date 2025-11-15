package dataForm

import "lesson04/core/dao/mySQL/models"

type GetSelectableClassReply struct {
	Action  string                     `json:"action"`
	Status  string                     `json:"status"`
	Message string                     `json:"message"`
	Classes []models.SelectableClasses `json:"classes"`
}

type GetStudentsClassReply struct {
	Action  string               `json:"action"`
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Classes models.StudentsTable `json:"classes"`
}

type DelStudentsClassReply struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AddStudentsClassForm struct {
	Action  string `json:"action"`
	ClassID string `json:"class_id"`
}

type AddStudentsClassReply struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
