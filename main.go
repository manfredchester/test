package main

import (
	"test/zhlog"
)

func main() {
	AzureBillTrans()
}

func AzureBillTrans() {
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error("AzureBillTrans", "%s", e.(error).Error())
		}
	}()
	InitCsvToDbByMonth(121)
	TransMasterByMonth(121)
	TransTagByMonth(121)
}
