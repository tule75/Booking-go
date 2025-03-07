package initialize

import (
	"database/sql"
	"ecommerce_go/global"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQLC() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := sql.Open("mysql", dsn)
	checkErrorPanic(err, "Initial mysql connection error")

	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdbc = db
	SetPoolC()
}

func SetPoolC() {
	sqlDb := global.Mdbc

	sqlDb.SetConnMaxIdleTime(time.Duration(global.Config.Mysql.MaxIdleConns))
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifeTime))

}
