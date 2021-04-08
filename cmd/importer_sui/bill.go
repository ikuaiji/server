package main

import (
	"app"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//随手记可能有多种时间记录格式
var timeLayouts = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
}

var billTypeMap = map[string]string{
	"转账":   "transfer",
	"支出":   "outcome",
	"收入":   "income",
	"余额变更": "alter",
}

//固定为东八区
var tz = time.FixedZone("CST", 8*60*60)

func parseTime(s string) (t time.Time, err error) {
	for _, layout := range timeLayouts {
		t, err = time.ParseInLocation(layout, s, tz)
		if err == nil {
			return
		}
	}

	return
}

func parseBill(file *excelize.File, categoryIdMap, accountIdMap, memberIdMap, projectIdMap, storeIdMap map[string]uint) ([]*app.Bill, error) {
	var bills []*app.Bill

	for _, sheet := range sheets {
		rows, err := file.GetRows(sheet)
		if err != nil {
			return nil, err
		}

		for _, row := range rows[1:] {
			if row[0] == "" {
				continue
			}

			amount, err := strconv.ParseFloat(row[5], 32)
			if err != nil {
				return nil, err
			}

			createdAt, err := parseTime(row[6])
			if err != nil {
				return nil, err
			}

			bill := &app.Bill{
				Type:       billTypeMap[row[0]],
				CategoryId: categoryIdMap[row[2]],
				AccountId:  accountIdMap[row[3]],
				Account2Id: accountIdMap[row[4]],
				Amount:     float32(amount),
				BillAt:     createdAt,
				MemberId:   memberIdMap[row[7]],
				ProjectId:  projectIdMap[row[8]],
				StoreId:    storeIdMap[row[9]],
				Note:       row[10],
			}

			bills = append(bills, bill)
		}

	}
	return bills, nil
}
