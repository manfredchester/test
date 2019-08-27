package main

import (
	"api-service-platform/types/azureType"
	"api-service-platform/util/cloudBillUtil"
	"api-service-platform/zhlog"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"test/common"
	"time"
)

func GetLoops(len int) (times int) {
	times = len / common.MAXINSERTNUM
	if len%common.MAXINSERTNUM != 0 {
		times++
	}
	return
}

func ReadCsv(filename string) [][]string {
	f, err := os.Open(filename)
	zhlog.Assert(err)
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	f.Close()
	zhlog.Assert(err)
	return rows
}

func GetCycle(filename string) string {
	cyc := strings.Split(filename, "_")
	billCycle := strings.Split(cyc[2], ".")
	return billCycle[0]
}

func InitCsvToDbByMonth(accountID int) {
	orm := CloudprojectEngine()
	orm.ShowSQL(false)
	session := orm.NewSession()
	defer session.Close()
	err := session.Begin()
	zhlog.Assert(err)

	v := fmt.Sprintf("dataCsv/bmw.csv")
	zhlog.Log("InitCsvToDbByMonth_files", "files: %s", v)

	data := ReadCsv(v)
	zhlog.Log("InitCsvToDbByMonth_num:", "len(data)：%d", len(data))
	zhlog.Log("InitCsvToDbByMonth_times:", "GetLoops(len(data))：%d", GetLoops(len(data)))

	var affectedAll int64
	for t := 0; t < cloudBillUtil.GetLoops(len(data)); t++ {
		// for t := 0; t < 2; t++ {
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
		affected, err := session.Insert(&ssd)
		affectedAll = affectedAll + affected
		if err != nil {
			session.Rollback()
			zhlog.Assert(errors.New(fmt.Sprintf("%d至%d发生错误:", t*common.MAXINSERTNUM, (t+1)*common.MAXINSERTNUM)))
		}
		session.Commit()
	}
	zhlog.Log("InitCsvToDbByMonth", "InitCsvToDbByMonth_affected: %d", affectedAll)

	fmt.Println("---------------------InitCsvToDbByMonth-------------end-------------------")
}

func TransMasterByMonth(accountID int) {
	orm := CloudprojectEngine()
	orm.ShowSQL(false)
	session := orm.NewSession()
	defer session.Close()
	err := session.Begin()
	zhlog.Assert(err)

	// accountCloudID, err := GetAccountCloudId(session, accountID)
	// zhlog.Assert(err)

	var data azureType.BillAzureDetailReportMetadata
	num, err := session.Where("bill_account_uuid=?", accountID).And("deleted=0").Count(data)
	// num, err := session.Where("bill_account_uuid=?", accountID).And("deleted=0").And("billing_cycle=?", curMonth).Count(data)

	zhlog.Assert(err)
	zhlog.Log("TransMasterByMonth", "TransMasterByMonth_num: %d", num)

	times := num / common.DefaultRowsLimit
	if num%common.DefaultRowsLimit != 0 {
		times++
	}

	zhlog.Log("TransMasterByMonth", "TransMasterByMonth_times: %d", times)

	var affectedAll int64
	for i := 0; i < int(times); i++ {
		var dataTo []azureType.BillAzureDetailReportMetadata
		var transData []azureType.BillCloudMaster
		// err = session.Where("bill_account_uuid=?", accountID).And("deleted=0").And("billing_cycle=?", curMonth).OrderBy("date").Limit(common.DefaultRowsLimit, common.DefaultRowsLimit*i).Find(&dataTo)
		err = session.Where("bill_account_uuid=?", accountID).And("deleted=0").OrderBy("date").Limit(common.DefaultRowsLimit, common.DefaultRowsLimit*i).Find(&dataTo)
		zhlog.NotNilErrorAssert("Begin:", err)

		for _, v := range dataTo {
			transData = append(transData, azureType.BillCloudMaster{
				ItemUuid:        v.Uuid,
				BillAccountUuid: v.BillAccountUuid,
				CloudType:       "azure",
				AccountId:       "V5701903S0105",
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

		affected, err := session.Insert(&transData)
		affectedAll = affectedAll + affected
		if zhlog.IsNotNil(err) {
			session.Rollback()
		}
		session.Commit()

	}
	zhlog.Log("TransMasterByMonth", "TransMasterByMonth_affected: %d", affectedAll)
	fmt.Println("---------------------TransMasterByMonth-------------end-------------------")
}

func TransTagByMonth(accountID int) {
	orm := CloudprojectEngine()
	orm.ShowSQL(false)
	session := orm.NewSession()
	defer session.Close()
	err := session.Begin()
	zhlog.Assert(err)

	var master azureType.BillCloudMaster
	// num, err := session.Where("bill_account_uuid=?", accountID).And("billing_cycle=?", curMonth).And("deleted=0").Count(master)
	num, err := session.Where("bill_account_uuid=?", accountID).And("deleted=0").Count(master)

	zhlog.Log("TransTagByMonth", "TransTagByMonth_num: %d", num)
	zhlog.Assert(err)

	times := num / common.DefaultRowsLimitTAG
	if num%common.DefaultRowsLimitTAG != 0 {
		times++
	}
	zhlog.Log("TransTagByMonth", "TransTagByMonth_times: %d", times)

	var affectedAll int64
	for i := 0; i < int(times); i++ {
		var masterTo []azureType.BillCloudMaster
		var transTag []azureType.BillCloudTag
		// err = session.Where("bill_account_uuid=?", accountID).And("billing_cycle=?", curMonth).And("deleted=0").OrderBy("date").Limit(common.DefaultRowsLimitTAG, common.DefaultRowsLimitTAG*i).Find(&masterTo)
		err = session.Where("bill_account_uuid=?", accountID).And("deleted=0").OrderBy("date").Limit(common.DefaultRowsLimitTAG, common.DefaultRowsLimitTAG*i).Find(&masterTo)

		zhlog.NotNilErrorAssert("Begin:", err)

		for _, v := range masterTo {
			if v.Tags != "" {
				jMap := make(map[string]interface{}, 0)
				err := json.Unmarshal([]byte(v.Tags), &jMap)
				zhlog.Assert(err)
				transTag = Analy(jMap, v.ItemUuid, v.BillAccountUuid, v.BillingCycle, v.AccountId)
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

		affected, err := session.Insert(&transTag)
		affectedAll = affectedAll + affected
		if zhlog.IsNotNil(err) {
			fmt.Println("err:", err)
			session.Rollback()
		}
		session.Commit()
	}
	zhlog.Log("TransTagByMonth", "TransTagByMonth_affected: %d", affectedAll)
	fmt.Println("---------------------TransTagByMonth-------------end-------------------")
}
