package dataForm

import (
	"lesson04/core/dao/mySQL/models"
)

type RegisterStudentsForm struct {
	Action string               `json:"action"`
	Data   models.StudentsTable `json:"data"`
}

type RegisterStudentsReply struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginStudentsForm struct {
	Action    string `json:"action"`
	StudentID string `json:"student_id"`
	Password  string `json:"password"`
}

type LoginStudentsReply struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type DelStudentsReply struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type StudentsInfoReply struct {
	Action  string               `json:"action"`
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Info    models.StudentsTable `json:"info"`
}

type StudentsUpdateForm struct {
	Action string  `json:"action"`
	Data   SubForm `json:"data"`
}

type SubForm struct {
	StudentPasswd string `json:"student_passwd"`
	StudentName   string `json:"student_name"`
	Sex           int    `json:"sex"`
	Grade         string `json:"grade"`
}

type StudentsUpdateReply struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
