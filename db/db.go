package db

import (
	"app"

	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

//Init初始化数据库连接
func Init(dsn string) error {
	//强制设置几个重要的DSN Params
	dbConfig, err := sqldriver.ParseDSN(dsn)
	if err != nil {
		return err
	}
	if dbConfig.Params == nil {
		dbConfig.Params = make(map[string]string)
	}
	dbConfig.Params["charset"] = "utf8mb4"
	dbConfig.Params["parseTime"] = "True"

	db, err := gorm.Open(mysql.Open(dbConfig.FormatDSN()), &gorm.Config{})
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
