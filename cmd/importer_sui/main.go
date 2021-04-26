package main

import (
	"app"
	"app/db"
	"flag"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var truncateAll bool
var dsn string
var filename = "myMoney.xlsx"
var sheets = []string{"支出", "收入", "余额变更", "转账"}

func init() {
	flag.BoolVar(&truncateAll, "truncate", true, "是否清空现有数据")
	flag.StringVar(&dsn, "dsn", "root:root@tcp(127.0.0.1:3306)/ikuaiji?charset=utf8mb4&parseTime=True&loc=Local", "数据库DSN")
}

func main() {
	flag.Parse()

	err := db.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	metaValueMaps, err := parseMetaInfo(f)
	if err != nil {
		log.Fatal(err)
	}

	if truncateAll {
		db.TruncateAllTable()
	}

	members := make([]*app.Member, 0, len(metaValueMaps["成员"]))
	for name := range metaValueMaps["成员"] {
		members = append(members, &app.Member{Name: name})
	}
	db.Save(&members)
	memberIdMap := make(map[string]uint, len(members))
	for _, member := range members {
		memberIdMap[member.Name] = member.ID
	}

	projects := make([]*app.Project, 0, len(metaValueMaps["项目"]))
	for name := range metaValueMaps["项目"] {
		projects = append(projects, &app.Project{Name: name})
	}
	db.Save(&projects)
	projectIdMap := make(map[string]uint, len(projects))
	for _, project := range projects {
		projectIdMap[project.Name] = project.ID
	}

	stores := make([]*app.Store, 0, len(metaValueMaps["商家"]))
	for name := range metaValueMaps["商家"] {
		stores = append(stores, &app.Store{Name: name})
	}
	db.Save(&stores)
	storeIdMap := make(map[string]uint, len(stores))
	for _, store := range stores {
		storeIdMap[store.Name] = store.ID
	}

	accounts := make([]*app.Account, 0, len(metaValueMaps["账户"]))
	for name := range metaValueMaps["账户"] {
		accounts = append(accounts, &app.Account{Name: name})
	}
	db.Save(&accounts)
	accountIdMap := make(map[string]uint, len(accounts))
	for _, account := range accounts {
		accountIdMap[account.Name] = account.ID
	}

	//先保存主分类
	mainCategories := make([]*app.Category, 0, len(metaValueMaps["主分类"]))
	for name := range metaValueMaps["主分类"] {
		mainCategories = append(mainCategories, &app.Category{Name: name})
	}
	db.Save(&mainCategories)

	mainCategoryIdMap := map[string]uint{}
	for _, category := range mainCategories {
		mainCategoryIdMap[category.Name] = category.ID
	}

	subCategories := make([]*app.Category, 0, len(metaValueMaps["子分类"]))
	for name, parentName := range metaValueMaps["子分类"] {
		subCategories = append(subCategories, &app.Category{Name: name, ParentId: mainCategoryIdMap[parentName]})
	}
	db.Save(&subCategories)
	subCategoryIdMap := make(map[string]uint, len(subCategories))
	for _, subCategory := range subCategories {
		subCategoryIdMap[subCategory.Name] = subCategory.ID
	}

	bills, err := parseBill(f, subCategoryIdMap, accountIdMap, memberIdMap, projectIdMap, storeIdMap)
	if err != nil {
		log.Fatal(err)
	}
	db.Save(&bills)
}
