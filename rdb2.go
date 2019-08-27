package main

import (
	"api-service-platform/plugin/logs"
	"api-service-platform/types/azureType"
	"api-service-platform/zhlog"
	"errors"
	protoConfig "protocol/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Analy(obj interface{}, itemUuid string, billAccountUuid int, billingCycle, accountID string) (transTag []azureType.BillCloudTag) {
	switch obj.(type) {
	case map[string]interface{}:
		for k, v := range obj.(map[string]interface{}) {
			transTag = append(transTag, azureType.BillCloudTag{
				Uuid:            zhlog.UUID(128),
				ItemUuid:        itemUuid,
				BillAccountUuid: billAccountUuid,
				BillingCycle:    billingCycle,
				AccountId:       accountID,
				TagKey:          k,
				TagValue:        v.(string),
				TagType:         "Tags",
				Deleted:         0,
				CreateTime:      time.Now(),
				UpdateTime:      time.Now(),
			})
		}
	default:
		zhlog.Log("TagsAnaly", "%s", "ERROR：存在异常数据结构类型，并未清晰入cloud_tag")
	}
	return transTag
}

var orm *xorm.Engine

func CloudprojectEngine() *xorm.Engine {
	orm, err := GetEngine()
	if err != nil {
		logs.Error(err)
	}
	return orm
}

func GetEngine() (*xorm.Engine, error) {
	if orm == nil {
		var err error
		orm, err = mysqlEngine()
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
	}
	if orm == nil {
		err := errors.New("database init error")
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
	}

	orm.ShowSQL()
	return orm, nil
}

func mysqlEngine() (*xorm.Engine, error) {
	Host := protoConfig.CMysqlHost()
	Host = "10.128.0.180"
	Port := "3306"
	Name := "cloudproject"
	User := "root"
	Password := "Connext@0101"
	dburl := User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Name + "?charset=utf8"
	logs.Info("dburl:", dburl)
	return xorm.NewEngine("mysql", dburl)
	// return xorm.NewEngine("mysql", "connextpaas:connext@0101@tcp(127.0.0.1:3306)/vmwareproject?charset=utf8")
}
