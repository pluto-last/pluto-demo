package main

import (
	"gV2/global"
	"gV2/model/table"
	"log"
)

func batchInsert() {

	var mckeord []table.McKeyword

	for i := 0; i < 10000; i++ {
		temp := table.McKeyword{
			CityName: "1",
			Keyword:  "1",
			Status:   "1",
		}

		mckeord = append(mckeord, temp)
	}

	// 制定一批插入的数据
	err := global.GVA_DB.CreateInBatches(mckeord, 1000).Error

	if err == nil {
		log.Println("批量插入成功")
	}

}

func findInMap() {
	var results []map[string]interface{}
	global.GVA_DB.Table("mc_keyword").Find(&results)

	log.Println(results)
}
