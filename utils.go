package main

import (
	"test/azureType"
	"test/zhlog"
	"time"
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
