package core

import (
	llog "lesson04/core/utils/LocalLog"

	"gorm.io/gorm"
)

type Config struct {
	MySQLAddress  string `json:"MySQLAddress"`
	MySQLPort     string `json:"MySQLPort"`
	MySQLUser     string `json:"MySQLUser"`
	MySQLPassword string `json:"MySQLPassword"`
	MySQLDBName   string `json:"MySQLDBName"`
}

var (
	Logger *llog.Log
	DB     *gorm.DB
	Conf   Config
)

var (
	StatusOK    int = 200
	ServerError int = 501
	NotFound    int = 404
)
