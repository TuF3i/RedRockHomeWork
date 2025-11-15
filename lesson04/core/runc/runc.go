package runc

import (
	"lesson04/core"
	"lesson04/core/dao/mySQL/engine"
	llog "lesson04/core/utils/LocalLog"
	"lesson04/core/utils/config"
	"log"
	"time"

	"gitee.com/liumou_site/logger"
	"github.com/fatih/color"
)

func InitEnv() {
	logs := logger.NewLogger(1)
	logs.Modular = "Runc"

	runc := Runc{l: logs}
	runc.InitLogger("data/log")
}

func (root *Runc) InitLogger(path string) {
	root.l.Debug("Started to load mod<LocalLog>")

	flog := log.Logger{}
	flog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flog.SetPrefix("[FileLog] ")

	clog := logger.NewLogger(1)

	date := time.Now().Format("20060102")

	core.Logger = &llog.Log{
		LogPath:  path,
		LogLevel: "DEBUG",
		CLog:     clog,
		FLog:     &flog,
		ToDay:    date,
	}

	root.l.Info("Finished loading mod<LocalLog>")
	color.New(color.FgCyan).Add(color.Underline).Println("Present LocalLog Conf: ")
	c := color.New(color.FgYellow, color.Bold)
	c.Printf("\tLogPath: %v\n", path)
	c.Printf("\tLogLevel: %v\n", "DEBUG")
	c.Printf("\tCLog: %v\n", clog)
	c.Printf("\tFLog: %v\n", &flog)
	c.Printf("\tToDay: %v\n", date)

}

func (root *Runc) InitConfig(path string) {
	root.l.Debug("Started to load mod<Config>")
	if err := config.Init().GetConfig(path); err != nil {
		root.l.Warn("Load mod<Config> Error: %v", err.Error())
	}
	root.l.Info("Finished loading mod<LocalLog>")
	color.New(color.FgCyan).Add(color.Underline).Println("Present global config: ")
	c := color.New(color.FgYellow, color.Bold)
	c.Printf("\tMySQLAddress: %v\n", core.Conf.MySQLAddress)
	c.Printf("\tMySQLPort: %v\n", core.Conf.MySQLPort)
	c.Printf("\tMySQLUser: %v\n", core.Conf.MySQLUser)
	c.Printf("\tMySQLPassword: %v\n", core.Conf.MySQLPassword)
	c.Printf("\tMySQLDBName: %v\n", core.Conf.MySQLDBName)

}

func (root *Runc) InitMySQL() {
	root.l.Debug("Started to load mod<MySQL>")
	if err := engine.Init().InitMySQL(); err != nil {
		root.l.Warn("Load mod<MySQL> Error: %v", err.Error())
	}
	root.l.Info("Finished loading mod<MySQl>")
	color.New(color.FgCyan, color.Bold).Printf("\tGet MySQL Connection")
	color.New(color.FgGreen).Printf(" %v:%v ", core.Conf.MySQLAddress, core.Conf.MySQLPort)
	color.New(color.FgCyan, color.Bold).Printf("at position ")
	color.New(color.FgYellow).Add(color.Underline).Println(core.DB)
}
