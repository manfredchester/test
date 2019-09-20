package main

import (
	"encoding/json"
	"test/azureType"
	"test/common"
	"test/zhlog"
	"time"
)

func InitCsvToDbByMonth(accountID, fileName string) (interrupt bool) {
	zhlog.Log("BillAccountUuid", "BillAccountUuid: %+v", accountID)

	zhlog.Log("InitCsvToDbByMonth_files", "files: %s", fileName)
	data := ReadCsv(fileName)
	zhlog.Log("InitCsvToDbByMonth_num", "len(data)：%d", len(data))
	zhlog.Log("InitCsvToDbByMonth_times", "GetLoops(len(data))：%d", GetLoops(len(data)))

	orm := CloudprojectEngine()
	orm.ShowSQL(false)
	var affectedAll int64
	for t := 0; t < GetLoops(len(data)); t++ {
		var ssd []azureType.BillAzureDetailReportMetadata
		for i := t * common.MAXINSERTNUM; i < (t+1)*common.MAXINSERTNUM; i++ {
			if i < 3 || i >= len(data) {
				continue
			}
			ssd = append(ssd, azureType.BillAzureDetailReportMetadata{
				Uuid:            zhlog.UUID(128),
				BillAccountUuid: accountID,
				BillingCycle:    data[i][azureType.Date],

				AccountOwnerId:         data[i][azureType.AccountOwnerId],
				AccountName:            data[i][azureType.AccountName],
				ServiceAdministratorId: data[i][azureType.ServiceAdministratorId],
				SubscriptionId:         data[i][azureType.SubscriptionId],
				SubscriptionGuid:       data[i][azureType.SubscriptionGuid],
				SubscriptionName:       data[i][azureType.SubscriptionName],
				Date:                   data[i][azureType.Date],
				Month:                  data[i][azureType.Month],
				Day:                    data[i][azureType.Day],
				Year:                   data[i][azureType.Year],
				Product:                data[i][azureType.Product],
				MeterId:                data[i][azureType.MeterId],
				MeterCategory:          data[i][azureType.MeterCategory],
				MeterSubCategory:       data[i][azureType.MeterSubCategory],
				MeterRegion:            data[i][azureType.MeterRegion],
				MeterName:              data[i][azureType.MeterName],
				ConsumedQuantity:       data[i][azureType.ConsumedQuantity],
				ResourceRate:           data[i][azureType.ResourceRate],
				ExtendedCost:           data[i][azureType.ExtendedCost],
				ResourceLocation:       data[i][azureType.ResourceLocation],
				ConsumedService:        data[i][azureType.ConsumedService],
				InstanceId:             data[i][azureType.InstanceId],
				ServiceInfo1:           data[i][azureType.ServiceInfo1],
				ServiceInfo2:           data[i][azureType.ServiceInfo2],
				AdditionalInfo:         data[i][azureType.AdditionalInfo],
				Tags:                   data[i][azureType.Tags],
				StoreServiceIdentifier: data[i][azureType.StoreServiceIdentifier],
				DepartmentName:         data[i][azureType.DepartmentName],
				CostCenter:             data[i][azureType.CostCenter],
				UnitOfMeasure:          data[i][azureType.UnitOfMeasure],
				ResourceGroup:          data[i][azureType.ResourceGroup],

				Deleted:    0,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			})
		}
		affected, err := orm.Insert(&ssd)
		affectedAll = affectedAll + affected
		if err != nil {
			zhlog.Error("InitCsvToDbByMonth_Insert", "%d至%d发生错误: %v", t*common.MAXINSERTNUM, (t+1)*common.MAXINSERTNUM, err)
			interrupt = true
		}
	}
	zhlog.Log("InitCsvToDbByMonth", "InitCsvToDbByMonth_affected: %d \n ---------------------InitCsvToDbByMonth-------------end-------------------", affectedAll)
	return
}

func TransMasterByMonth(accountID, cldType, cldID string) (interrupt bool) {
	var data azureType.BillAzureDetailReportMetadata
	num, err := orm.Where("bill_account_uuid=?", accountID).And("deleted=0").Count(data)
	zhlog.Log("TransMasterByMonth", "TransMasterByMonth_num: %d", num)
	zhlog.Assert(err)

	times := num / common.DefaultRowsLimit
	if num%common.DefaultRowsLimit != 0 {
		times++
	}
	zhlog.Log("TransMasterByMonth", "TransMasterByMonth_times: %d", times)

	orm := CloudprojectEngine()
	orm.ShowSQL(false)
	var affectedAll int64
	for i := 0; i < int(times); i++ {
		dataTo := make([]azureType.BillAzureDetailReportMetadata, 0)
		transData := make([]azureType.BillCloudMaster, 0)
		zhlog.Log("Master执行状态", "Master执行数据 %+v Master开始值：%+v", common.DefaultRowsLimit, common.DefaultRowsLimit*i)
		err = orm.Where("bill_account_uuid=?", accountID).And("deleted=0").Limit(common.DefaultRowsLimit, common.DefaultRowsLimit*i).Find(&dataTo)
		zhlog.Assert(err)

		for _, v := range dataTo {
			transData = append(transData, azureType.BillCloudMaster{
				ItemUuid:        v.Uuid,
				BillAccountUuid: v.BillAccountUuid,
				CloudType:       cldType,
				AccountId:       cldID,
				BillingCycle:    v.BillingCycle,

				Date:             v.Date,
				SubscriptionGuid: v.SubscriptionGuid,
				SubscriptionName: v.SubscriptionName,
				Product:          v.Product,
				MeterCategory:    v.MeterCategory,
				ConsumedQuantity: v.ConsumedQuantity,
				ExtendedCost:     v.ExtendedCost,
				Tags:             v.Tags,
				ResourceGroup:    v.ResourceGroup,

				Deleted:    0,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			})
		}

		affected, err := orm.Insert(&transData)
		affectedAll = affectedAll + affected
		if zhlog.IsNotNil(err) {
			zhlog.Error("TransMasterByMonth_Insert", "%d至%d发生错误: %v", i*common.DefaultRowsLimit, (i+1)*common.DefaultRowsLimit, err)
			interrupt = true
		}
	}
	zhlog.Log("TransMasterByMonth", "TransMasterByMonth_affected: %d \n---------------------TransMasterByMonth-------------end-------------------", affectedAll)
	return
}

func TransTagByMonth(accountID string) {
	var master azureType.BillCloudMaster
	num, err := orm.Where("bill_account_uuid=?", accountID).And("deleted=0").Count(master)
	zhlog.Log("TransTagByMonth", "TransTagByMonth_num: %d", num)
	zhlog.Assert(err)

	times := num / common.DefaultRowsLimitTAG
	if num%common.DefaultRowsLimitTAG != 0 {
		times++
	}
	zhlog.Log("TransTagByMonth", "TransTagByMonth_times: %d", times)

	orm := CloudprojectEngine()
	orm.ShowSQL(false)
	var affectedAll int64
	for i := 0; i < int(times); i++ {
		masterTo := make([]azureType.BillCloudMaster, 0)
		transTag := make([]azureType.BillCloudTag, 0)
		zhlog.Log("TAG执行状态", "TAG执行数据 %+v TAG开始值：%+v", common.DefaultRowsLimit, common.DefaultRowsLimit*i)
		err = orm.Where("bill_account_uuid=?", accountID).And("deleted=0").OrderBy("item_uuid").Limit(common.DefaultRowsLimitTAG, common.DefaultRowsLimitTAG*i).Find(&masterTo)
		zhlog.Assert(err)

		for _, v := range masterTo {
			if v.Tags != "" {
				jMap := make(map[string]interface{}, 0)
				err := json.Unmarshal([]byte(v.Tags), &jMap)
				zhlog.Assert(err)
				transTag = append(transTag, Analy(jMap, v.ItemUuid, v.BillAccountUuid, v.BillingCycle, v.AccountId)...)
			}
			if v.SubscriptionGuid != "" {
				transTag = append(transTag, azureType.BillCloudTag{
					Uuid: zhlog.UUID(128),

					ItemUuid:        v.ItemUuid,
					BillAccountUuid: v.BillAccountUuid,

					BillingCycle: v.BillingCycle,
					AccountId:    v.AccountId,
					TagKey:       "",
					TagValue:     v.SubscriptionGuid,
					TagType:      "SubscriptionGuid",

					Deleted:    0,
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				})
			}
			if v.MeterCategory != "" {
				transTag = append(transTag, azureType.BillCloudTag{
					Uuid: zhlog.UUID(128),

					ItemUuid:        v.ItemUuid,
					BillAccountUuid: v.BillAccountUuid,

					BillingCycle: v.BillingCycle,
					AccountId:    v.AccountId,
					TagKey:       "",
					TagValue:     v.MeterCategory,
					TagType:      "MeterCategory",

					Deleted:    0,
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				})
			}
			if v.ResourceGroup != "" {
				transTag = append(transTag, azureType.BillCloudTag{
					Uuid: zhlog.UUID(128),

					ItemUuid:        v.ItemUuid,
					BillAccountUuid: v.BillAccountUuid,

					BillingCycle: v.BillingCycle,
					AccountId:    v.AccountId,
					TagKey:       "",
					TagValue:     v.ResourceGroup,
					TagType:      "ResourceGroup",

					Deleted:    0,
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				})
			}
		}

		affected, err := orm.Insert(&transTag)
		affectedAll = affectedAll + affected
		if zhlog.IsNotNil(err) {
			zhlog.Error("TransTagByMonth_affected", "%d至%d发生错误: %v ", i*common.DefaultRowsLimitTAG, (i+1)*common.DefaultRowsLimitTAG, err)
		}
	}
	zhlog.Log("TransTagByMonth", "TransTagByMonth_affected: %d \n ---------------------TransTagByMonth-------------end-------------------", affectedAll)
}
