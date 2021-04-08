package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//metaColTitleRemap保存所有需要提取为meta的表头
var metaColTitleRemap = map[string]string{
	//"交易类型":SKIP
	"分类":  "主分类",
	"子分类": "子分类",
	"账户1": "账户",
	"账户2": "账户",
	//"金额":SKIP
	//"日期":SKIP
	"成员": "成员",
	"项目": "项目",
	"商家": "商家",
	//"备注":SKIP
}

//parseMetaInfo 从Excel表中解析记账所需的meta信息，返回各个信息的map
//每个信息map的key为该meta名，value除子分类为起主分类名外，其他为空字符串
func parseMetaInfo(file *excelize.File) (map[string]map[string]string, error) {
	valueMaps := map[string]map[string]string{
		"主分类": map[string]string{},
		"子分类": map[string]string{},
		"账户":  map[string]string{},
		"成员":  map[string]string{},
		"项目":  map[string]string{},
		"商家":  map[string]string{},
	}

	//分类和子分类
	var mainCategoryRows, subCategoryRows []string
	for _, sheet := range sheets {
		cols, err := file.GetCols(sheet)
		if err != nil {
			return nil, err
		}
		for _, col := range cols {
			title, ok := metaColTitleRemap[col[0]]
			if !ok {
				continue
			}

			for _, val := range col[1:] {
				if val != "" {
					valueMaps[title][val] = ""
				}
			}

			if col[0] == "分类" {
				mainCategoryRows = append(mainCategoryRows, col[1:]...)
			}

			if col[0] == "子分类" {
				subCategoryRows = append(subCategoryRows, col[1:]...)
			}
		}
	}

	for i, subCategoryRow := range subCategoryRows {
		if subCategoryRow != "" {
			valueMaps["子分类"][subCategoryRow] = mainCategoryRows[i]
		}
	}

	return valueMaps, nil
}
