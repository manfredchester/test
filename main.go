package main

import (
	"api-service-platform/zhlog"
)

func main() {
	AzureBillTrans()
}

// func logTaketime(startTime int64) int {
// 	endTimd := time.Now().UnixNano() / 1e6
// 	logtime := int(endTimd - startTime)
// 	return logtime
// }

func AzureBillTrans() {
	// var curMonth string
	// start := time.Now().UnixNano() / 1e6
	defer func() {
		if e := recover(); e != nil {
			// rdb.UpdateAccountBillSync(accountID, curMonth, 1)
			// rdb.ChangeTask(taskID, types.TaskStateRunError)
			// zhlog.ErrorToDBPro("AzureBillTrans", strconv.Itoa(taskID), "CMDBLOG_AZURE_BILL", logTaketime(start), "ERROR：%s", e.(error).Error())
			zhlog.Error("AzureBillTrans", "%s", e.(error).Error())
		}
	}()
	// rdb.ChangeTask(taskID, types.TaskStateRunning)

	// enrollNum, apikey := rdb.GetAccountInfo(accountID)
	// if enrollNum == "" || apikey == "" {
	// 	zhlog.Assert(errors.New("账号信息为空，请检查..."))
	// }
	// decryptCode := AesDecrypt(apikey, common.EncryptionKey)
	// usage := GetURLList(enrollNum, decryptCode)

	// billAttr := rdb.GetBillAttr(accountID)
	// if len(billAttr) != 1 {
	// 	zhlog.Assert(errors.New("账号云账单初始化信息不存在或存在脏数据"))
	// }
	// initStartTime := billAttr[0].InitStartDate
	// if initStartTime == "" {
	// 	zhlog.Assert(errors.New("无法获取初始化时间，请在账户管理界面设置初始化时间"))
	// }
	// initEndTime := time.Now().Format("2006-01")

	// local, err := time.LoadLocation("Local")
	// zhlog.Assert(err)
	// str, err := time.ParseInLocation("2006-01", initStartTime, local)
	// zhlog.Assert(err)

	// for {
	// 	curMonth = str.Format("2006-01")
	// 	// 开始时间大于本月的时间，直接退出
	// 	if curMonth > initEndTime {
	// 		break
	// 	}
	// 	// 判断该时间若在数据库中是否存有成功的记录
	// 	if ext := rdb.QueryAccountBillSync(accountID, curMonth, 3); !ext {
	// 		// 成功记录不存在；不存在则需要更新或加入该时间数据
	// 		// 判断是否存在该时间，用来更新或拆入执行记录
	// 		if curMonth != initEndTime {
	// 			if ext := rdb.QueryAccountBillSync(accountID, curMonth, -1); ext {
	// 				// 存在，则是更新记录的执行状态
	// 				rdb.UpdateAccountBillSync(accountID, curMonth, 2)
	// 			} else {
	// 				// 不存在，则是拆入执行的记录信息
	// 				rdb.InsertAccountBillSync(accountID, curMonth, 2)
	// 			}
	// 		}

	// 		var cnt int
	// 		for i := 0; i < len(usage.AvailableMonths); i++ {
	// 			if usage.AvailableMonths[i].Month == curMonth {
	// 				cnt = i
	// 				break
	// 			}
	// 		}
	// 		zhlog.Log("AzureBillTrans---DailyUpdate", "accountID: %d\n ,enrollNum: %s\n ,curMonth: %s\n ,cnt: %d \n", accountID, enrollNum, curMonth, cnt)

	// 		fileName := fmt.Sprintf("dataCsv/%s_%s.csv", enrollNum, curMonth)
	// 		os.Remove(fileName)
	// 		// err := os.Remove(fileName)
	// 		// zhlog.Assert(err)

	// 		// files, err := filepath.Glob("dataCsv/*")
	// 		// zhlog.Assert(err)
	// 		// zhlog.Log("DailyUpdateRemove", "files: %s", files)

	// 		GetData(enrollNum, usage, apikey, cnt, accountID)

	// files, err = filepath.Glob("dataCsv/*")
	// zhlog.Assert(err)
	// zhlog.Log("DailyUpdateGetData", "files: %s", files)

	// ClearDatabyMonth(accountID, curMonth)
	InitCsvToDbByMonth(88)
	TransMasterByMonth(88)
	TransTagByMonth(88)
	// 如果是当前月 则不会记录操作
	// 		if curMonth != initEndTime {
	// 			rdb.UpdateAccountBillSync(accountID, curMonth, 3)
	// 		}
	// 	}

	// 	str = str.AddDate(0, 1, 0)
	// }
	// rdb.ChangeTask(taskID, types.TaskStateRunDone)
	// zhlog.InfoToDB("AzureBillTrans", strconv.Itoa(taskID), "CMDBLOG_AZURE_BILL", logTaketime(start), "Azure 企业账号： %s \n 云账单处理成功", enrollNum)
}
