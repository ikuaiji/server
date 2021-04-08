package db

import (
	"app"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

//Init初始化数据库连接
//dsn格式： user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func Init(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&app.Bill{}, &app.Category{}, &app.Account{}, &app.Member{}, &app.Project{}, &app.Store{})

	dbConn = db
	return nil
}

func TruncateAllTable() error {
	var err error

	err = dbConn.Migrator().DropTable(&app.Bill{}, &app.Category{}, &app.Account{}, &app.Member{}, &app.Project{}, &app.Store{})
	if err != nil {
		return err
	}

	err = dbConn.AutoMigrate(&app.Bill{}, &app.Category{}, &app.Account{}, &app.Member{}, &app.Project{}, &app.Store{})
	if err != nil {
		return err
	}

	return nil
}

func Save(data interface{}) *gorm.DB {
	return dbConn.Save(data)
}
