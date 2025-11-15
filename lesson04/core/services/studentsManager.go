package services

import (
	"fmt"
	"lesson04/core"
	"lesson04/core/dao/mySQL/handleFunc"
	"lesson04/core/dao/mySQL/models"
	"lesson04/core/services/dataForm"
	"lesson04/core/utils/jwt"
	"lesson04/core/utils/md5"
	"strings"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func (root *Service) StudentRegister() {
	var stuDataForm dataForm.RegisterStudentsForm
	mySQLHandle := handleFunc.Init()
	sourceIP := root.c.ClientIP()

	if err := root.c.BindAndValidate(&stuDataForm); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusBadRequest, dataForm.RegisterStudentsReply{
			Action:  stuDataForm.Action,
			Status:  "Failed",
			Message: "参数错误",
		})
		return
	}

	if mySQLHandle.CheckStudentExists(stuDataForm.Data.StudentID) {
		core.Logger.BotDEBUG(fmt.Sprintf("%v - %v", sourceIP, "Student already exists!"))
		root.c.JSON(consts.StatusConflict, dataForm.RegisterStudentsReply{
			Action:  stuDataForm.Action,
			Status:  "Illegal",
			Message: "学生已存在",
		})
		return
	}

	if err := mySQLHandle.AddStudent(stuDataForm.Data); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.RegisterStudentsReply{
			Action:  stuDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, stuDataForm.Action+" Success"))
	root.c.JSON(consts.StatusOK, dataForm.RegisterStudentsReply{
		Action:  stuDataForm.Action,
		Status:  "Success",
		Message: fmt.Sprintf("操作成功"),
	})

	//root.c.Redirect(consts.StatusFound, []byte(core.UrlLogin))

	return

}

func (root *Service) StudentLogin() {
	var revDataForm dataForm.LoginStudentsForm
	mySQLHandle := handleFunc.Init()
	sourceIP := root.c.ClientIP()

	if err := root.c.BindAndValidate(&revDataForm); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusBadRequest, dataForm.LoginStudentsReply{
			Action:  revDataForm.Action,
			Status:  "Failed",
			Message: "参数错误",
			Token:   "",
		})
		return
	}

	if mySQLHandle.CheckStudentExists(revDataForm.StudentID) {
		core.Logger.BotDEBUG(fmt.Sprintf("%v - %v", sourceIP, "Student not exists!"))
		root.c.JSON(consts.StatusConflict, dataForm.LoginStudentsReply{
			Action:  revDataForm.Action,
			Status:  "Illegal",
			Message: "学生不存在",
			Token:   "",
		})
		return
	}

	data, err := mySQLHandle.GetStudentInfo(revDataForm.StudentID)

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.LoginStudentsReply{
			Action:  revDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Token:   "",
		})
		return
	}

	if data.StudentPasswd != md5.GenMD5(revDataForm.Password) {
		core.Logger.BotDEBUG(fmt.Sprintf("%v - %v", sourceIP, "Wrong Password!"))
		root.c.JSON(consts.StatusUnauthorized, dataForm.LoginStudentsReply{
			Action:  revDataForm.Action,
			Status:  "illegal",
			Message: "密码错误",
			Token:   "",
		})
		return
	}

	jwtGen, err := jwt.InitJWT()
	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.LoginStudentsReply{
			Action:  revDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Token:   "",
		})
		return
	}

	token, err := jwtGen.GenJWT(data.StudentID)
	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.LoginStudentsReply{
			Action:  revDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Token:   "",
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, revDataForm.Action+" Success"))
	root.c.JSON(consts.StatusOK, dataForm.LoginStudentsReply{
		Action:  revDataForm.Action,
		Status:  "Success",
		Message: "登陆成功",
		Token:   token,
	})
	return
}

func (root *Service) StudentDel() {
	mySQLHandle := handleFunc.Init()
	jwtChecker, err := jwt.InitJWT()
	sourceIP := root.c.ClientIP()

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.DelStudentsReply{
			Action:  "DELETE STUDENT",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
		return
	}

	authHeader := root.c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentID, ok := jwtChecker.RecoverData(token)
	if !ok {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, "Decode JWT Error"))
		root.c.JSON(consts.StatusInternalServerError, dataForm.DelStudentsReply{
			Action:  "DELETE STUDENT",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
		})
		return
	}

	if !mySQLHandle.CheckStudentExists(studentID) {
		core.Logger.BotDEBUG(fmt.Sprintf("%v - %v", sourceIP, "Student not exists!"))
		root.c.JSON(consts.StatusConflict, dataForm.DelStudentsReply{
			Action:  "DELETE STUDENT",
			Status:  "Illegal",
			Message: "学生不存在",
		})
		return
	}

	if err := mySQLHandle.DeleteStudent(studentID); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.DelStudentsReply{
			Action:  "DELETE STUDENT",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, "DELETE STUDENT Success"))
	root.c.JSON(consts.StatusOK, dataForm.DelStudentsReply{
		Action:  "DELETE STUDENT",
		Status:  "Success",
		Message: "操作成功",
	})

	return
}

func (root *Service) StudentInfo() {
	mySQLHandle := handleFunc.Init()
	jwtChecker, err := jwt.InitJWT()
	sourceIP := root.c.ClientIP()
	var student models.StudentsTable

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsInfoReply{
			Action:  "GET STU INFO",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Info:    student,
		})
		return
	}

	authHeader := string(root.c.Request.Header.Get("Authorization"))
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentID, ok := jwtChecker.RecoverData(token)
	if !ok {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, "Decode JWT Error"))
		root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsInfoReply{
			Action:  "GET STU INFO",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
			Info:    student,
		})
		return
	}

	student, err = mySQLHandle.GetStudentInfo(studentID)
	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsInfoReply{
			Action:  "GET STU INFO",
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
			Info:    student,
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, "GET STU INFO Success"))
	root.c.JSON(consts.StatusOK, dataForm.StudentsInfoReply{
		Action:  "GET STU INFO",
		Status:  "Success",
		Message: "操作成功",
		Info:    student,
	})

	return
}

func (root *Service) StudentUpdate() {
	var updateDataForm dataForm.StudentsUpdateForm
	mySQLHandle := handleFunc.Init()
	jwtChecker, err := jwt.InitJWT()
	sourceIP := root.c.ClientIP()

	if err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsUpdateReply{
			Action:  updateDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
		return
	}

	if err := root.c.BindAndValidate(&updateDataForm); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusBadRequest, dataForm.StudentsUpdateReply{
			Action:  updateDataForm.Action,
			Status:  "Failed",
			Message: "参数错误",
		})
		return
	}

	authHeader := string(root.c.Request.Header.Get("Authorization"))
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentID, ok := jwtChecker.RecoverData(token)
	if !ok {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, "Decode JWT Error"))
		root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsUpdateReply{
			Action:  updateDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", "Decode JWT Error"),
		})
		return
	}

	data := models.StudentsTable{
		StudentID:     studentID,
		StudentPasswd: updateDataForm.Data.StudentPasswd,
		StudentName:   updateDataForm.Data.StudentName,
		Sex:           updateDataForm.Data.Sex,
		Grade:         updateDataForm.Data.Grade,
	}

	if err := mySQLHandle.UpdateStudent(data); err != nil {
		core.Logger.BotWarning(fmt.Sprintf("%v - %v", sourceIP, err.Error()))
		root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsUpdateReply{
			Action:  updateDataForm.Action,
			Status:  "Failed",
			Message: fmt.Sprintf("服务器内部错误：%v", err.Error()),
		})
		return
	}

	core.Logger.BotINFO(fmt.Sprintf("%v - %v", sourceIP, updateDataForm.Action+" Success"))
	root.c.JSON(consts.StatusInternalServerError, dataForm.StudentsUpdateReply{
		Action:  updateDataForm.Action,
		Status:  "Success",
		Message: "操作成功",
	})
	return
}
