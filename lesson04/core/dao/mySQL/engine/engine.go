package engine

import (
	"fmt"
	"lesson04/core"
	"lesson04/core/dao/mySQL/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var modelList = []interface{}{
	models.Link{},
	models.StudentsTable{},
	models.MainClassTable{},
	models.SelectableClasses{},
}

func (root *MySQL) InitMySQL() error {
	//构造MySQL URL
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		core.Conf.MySQLUser,
		core.Conf.MySQLPassword,
		core.Conf.MySQLAddress,
		core.Conf.MySQLPort,
		core.Conf.MySQLDBName)

	//连接数据库
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return err
	}

	//迁移数据库
	err = db.AutoMigrate(modelList...)
	if err != nil {
		return err
	}

	core.DB = db
	return nil
}
