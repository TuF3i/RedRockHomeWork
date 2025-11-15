package services

import (
	"fmt"
	"lesson04/core"
	"lesson04/core/dao/mySQL/handleFunc"
	"lesson04/core/dao/mySQL/models"
	"lesson04/core/services/dataForm"
	"lesson04/core/utils/jwt"
	"strings"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func (root *Service) GetSelectableClasses() {
	mySQLHandle := handleFunc.Init()
	sourceIP := root.c.ClientIP()

	classes, err := mySQLHandle.GetSelectableClasses()
	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.GetSelectableClassReply{
			Action:  "GET SELECTABLE CLASSES",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
			Classes: nil,
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, "GET SELECTABLE CLASSES Success"))
	root.c.JSON(consts.StatusOK, dataForm.GetSelectableClassReply{
		Action:  "GET SELECTABLE CLASSES",
		Status:  "Success",
		Message: "操作成功",
		Classes: classes,
	})
	return
}

func (root *Service) GetStudentsClasses() {
	mySQLHandle := handleFunc.Init()
	jwtChecker, err := jwt.InitJWT()
	sourceIP := root.c.ClientIP()

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.GetStudentsClassReply{
			Action:  "GET STUDENT CLASSES",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Classes: models.StudentsTable{},
		})
	}

	authHeader := string(root.c.Request.Header.Get("Authorization"))
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentID, ok := jwtChecker.RecoverData(token)
	if !ok {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, "Decode JWT Error"))
		root.c.JSON(consts.StatusInternalServerError, dataForm.GetStudentsClassReply{
			Action:  "GET STUDENT CLASSES",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
			Classes: models.StudentsTable{},
		})
		return
	}

	classes, err := mySQLHandle.GetClassesOfStu(studentID)

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.GetStudentsClassReply{
			Action:  "GET STUDENT CLASSES",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Classes: models.StudentsTable{},
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, "GET STUDENT CLASSES Success"))
	root.c.JSON(consts.StatusOK, dataForm.GetStudentsClassReply{
		Action:  "GET STUDENT CLASSES",
		Status:  "Success",
		Message: "操作成功",
		Classes: classes,
	})
	return
}

func (root *Service) DelStudentClass() {
	mySQLHandle := handleFunc.Init()
	jwtChecker, err := jwt.InitJWT()
	sourceIP := root.c.ClientIP()

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.DelStudentsClassReply{
			Action:  "DEL STUDENT CLASS",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
	}

	authHeader := string(root.c.Request.Header.Get("Authorization"))
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentID, ok := jwtChecker.RecoverData(token)
	if !ok {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, "Decode JWT Error"))
		root.c.JSON(consts.StatusInternalServerError, dataForm.DelStudentsClassReply{
			Action:  "DEL STUDENT CLASS",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
		})
		return
	}

	classID := root.c.Param("class_id")

	if mySQLHandle.CheckClassExists(studentID, classID) {
		core.Logger.BotDEBUG(fmt.Sprintf("%v - %v", sourceIP, "Class not exists!"))
		root.c.JSON(consts.StatusNotFound, dataForm.DelStudentsClassReply{
			Action:  "DEL STUDENT CLASS",
			Status:  "Illegal",
			Message: "课程不存在",
		})
		return
	}

	if err := mySQLHandle.DeleteClassForStu(studentID, classID); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.DelStudentsClassReply{
			Action:  "DEL STUDENT CLASS",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, "DEL STUDENT CLASS Success"))
	root.c.JSON(consts.StatusOK, dataForm.DelStudentsClassReply{
		Action:  "DEL STUDENT CLASS",
		Status:  "Success",
		Message: "操作成功",
	})
	return
}

func (root *Service) AddStudentClass() {
	var classData dataForm.AddStudentsClassForm
	mySQLHandle := handleFunc.Init()
	sourceIP := root.c.ClientIP()

	if err := root.c.BindAndValidate(&classData); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusBadRequest, dataForm.RegisterStudentsReply{
			Action:  classData.Action,
			Status:  "Failed",
			Message: "参数错误",
		})
		return
	}

	jwtChecker, err := jwt.InitJWT()
	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.AddStudentsClassReply{
			Action:  "ADD STUDENT CLASS",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
	}

	authHeader := string(root.c.Request.Header.Get("Authorization"))
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentID, ok := jwtChecker.RecoverData(token)
	if !ok {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, "Decode JWT Error"))
		root.c.JSON(consts.StatusInternalServerError, dataForm.AddStudentsClassReply{
			Action:  "ADD STUDENT CLASS",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
		})
		return
	}

	if !mySQLHandle.CheckClassExists(studentID, classData.ClassID) {
		core.Logger.BotDEBUG(fmt.Sprintf("%v - %v", sourceIP, "Class already exists!"))
		root.c.JSON(consts.StatusConflict, dataForm.AddStudentsClassReply{
			Action:  "ADD STUDENT CLASS",
			Status:  "Illegal",
			Message: "课程已存在",
		})
		return
	}

	if err := mySQLHandle.SelectClassForStu(studentID, classData.ClassID); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.AddStudentsClassReply{
			Action:  "ADD STUDENT CLASS",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, "ADD STUDENT CLASS Success"))
	root.c.JSON(consts.StatusOK, dataForm.AddStudentsClassReply{
		Action:  "ADD STUDENT CLASS",
		Status:  "Success",
		Message: "操作成功",
	})
	return

}
