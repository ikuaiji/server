package db

import (
	"app"
	"time"

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

//GetBillsOfMonth 获取指定月份的所有交易记录
func GetBillsOfMonth(year int, month time.Month) ([]app.Bill, error) {
	var records []app.Bill

	from := time.Date(year, month, 1, 0, 0, 0, 0, app.TZ)
	to := time.Date(year, month, 1, 0, 0, 0, 0, app.TZ).AddDate(0, 1, 0)

	result := dbConn.Where("bill_at BETWEEN ? AND ?", from, to).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

//GetMetaIdNameMap 获取所有Meta信息的id-name映射表，用于前端显示
func GetMetaIdNameMap() (map[string]map[uint]string, error) {
	type IdName struct {
		ID   uint
		Name string
	}

	tables := []string{
		"accounts",
		"categories",
		"members",
		"projects",
		"stores",
	}

	metaMap := make(map[string]map[uint]string)
	for _, table := range tables {
		var records []IdName
		result := dbConn.Table(table).Select("name", "id").Scan(&records)
		if result.Error != nil {
			return nil, result.Error
		}

		m := make(map[uint]string, len(records))
		for _, record := range records {
			m[record.ID] = record.Name
		}

		metaMap[table] = m
	}

	return metaMap, nil

}

//GetAccounts 获取所有账号的基础信息及余额
func GetAccounts() ([]app.Account, error) {
	var records []app.Account

	result := dbConn.Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

//GetAccountBalances 获取所有账号的余额，余额等于收入➖支出➕转入➖转出
func GetAccountBalances() (map[uint]float32, error) {
	balances := make(map[uint]float32)
	var income, outcome, trasferIn []app.Bill

	//先计算收入
	result := dbConn.Select("account_id, sum(amount) as amount").Where("type = 'income'").Group("account_id").Find(&income)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, record := range income {
		balances[record.AccountId] += record.Amount
	}

	//再计算支出和转账（转出）
	result = dbConn.Select("account_id, sum(amount) as amount").Where("type = 'outcome' OR type='transfer'").Group("account_id").Find(&outcome)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, record := range outcome {
		balances[record.AccountId] -= record.Amount
	}

	//再转账（转入）
	result = dbConn.Select("account2_id, sum(amount) as amount").Where("type='transfer'").Group("account2_id").Find(&trasferIn)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, record := range trasferIn {
		balances[record.Account2Id] += record.Amount
	}

	return balances, nil
}
