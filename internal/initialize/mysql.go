package initialize

import (
	"ecommerce_go/global"
	"ecommerce_go/internal/po"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErrorPanic(err, "Initial mysql connection error")

	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db
	SetPool()
	migrateTables()
}

func SetPool() {
	sqlDb, err := global.Mdb.DB()

	if err != nil {
		fmt.Println("msql error: ", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(global.Config.Mysql.MaxIdleConns))
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifeTime))

}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Println("Migration failed: ", err)
	}
}
