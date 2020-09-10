package azureType

import (
	"time"
)

// type BillAccountRsp struct {
// 	BillAccountUuid string `json:"bill_account_uuid"`
// 	AccountName     string `json:"account_name"`
// 	AccountId       string `json:"account_id"`
// 	CloudType       int    `json:"cloud_type"`
// 	InitDefaultTime string `json:"init_default_time"`
// 	InitStartTime   string `json:"init_start_time"`
// 	InitEndTime     string `json:"init_end_time"`
// 	Status          int    `json:"status"`
// }

// type BillAccount struct {
// 	BillAccountUuid string    `json:"bill_account_uuid"`
// 	AccountName     string    `json:"account_name"`
// 	AccountId       string    `json:"account_id"`
// 	CloudType       int       `json:"cloud_type"`
// 	AccessKey       string    `json:"access_key_id"`
// 	AccessSecret    string    `json:"access_secret"`
// 	InitDefaultTime string    `json:"init_default_time"`
// 	InitStartTime   string    `json:"init_start_time"`
// 	InitEndTime     string    `json:"init_end_time"`
// 	Status          int       `json:"status"`
// 	Deleted         int       `json:"deleted"`
// 	CreateTime      time.Time `json:"create_time"`
// 	UpdateTime      time.Time `json:"udpate_time"`
// }

// type BillAccountReq struct {
// 	BillAccountUuid string                     `json:"bill_account_uuid"`
// 	AccountName     string                     `json:"account_name"`
// 	AccountId       string                     `json:"account_id"`
// 	CloudType       int                        `json:"cloud_type"`
// 	AccessKey       string                     `json:"access_key_id"`
// 	AccessSecret    string                     `json:"access_secret"`
// 	InitStartTime   string                     `json:"init_start_time"`
// 	InitEndTime     string                     `json:"init_end_time"`
// 	Pagination      *request.RequestPagination `json:"pagination"`
// }

type BillAccountReqDemo struct {
	AccountInfo AccountInfo  `json:"account_info"`
	ClientInfo  []ClientInfo `json:"client_info"`
	// Pagination  *request.RequestPagination `json:"pagination"`
}

type AccountInfo struct {
	AccountId        int       `json:"account_id"`
	CloudType        string    `json:"cloud_type"`
	SyncTypeBill     string    `json:"sync_type_bill"`
	SyncTypeCmdb     string    `json:"sync_type_cmdb"`
	AccountCloudId   string    `json:"account_cloud_id"`
	AccountCloudName string    `json:"account_cloud_name"`
	AccountKeyId     string    `json:"account_key_id"`
	AccountSercet    string    `json:"account_secret"`
	TenantId         string    `json:"tenant_id"`
	CloudTypeName    string    `json:"cloud_type_name"`
	InitOldestDate   string    `json:"init_oldest_date"`
	InitStartDate    string    `json:"init_start_date"`
	UpdatedTime      time.Time `json:"updated_time"`
	HasLinkApp       string    `json:"has_link_app"` //是否连接应用
}

type ClientInfo struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type CloudAccount struct {
	AccountId        int       `json:"account_id"`
	AccountCloudId   string    `json:"account_cloud_id"`
	AccountCloudName string    `json:"account_cloud_name"`
	CloudType        string    `json:"cloud_type"`
	Enabled          int       `json:"enabled"`
	Deleted          int       `json:"deleted"`
	CreatedBy        string    `json:"created_by"`
	CreatedTime      time.Time `json:"created_time"`
	UpdatedBy        string    `json:"updated_by"`
	UpdatedTime      time.Time `json:"updated_time"`
}
type AccountList struct {
	AccountId        int       `json:"account_id"`
	AccountCloudId   string    `json:"account_cloud_id"`
	AccountCloudName string    `json:"account_cloud_name"`
	CloudType        string    `json:"cloud_type"`
	Enabled          int       `json:"enabled"`
	CloudTypeName    string    `json:"cloud_type_name"`
	InitOldestDate   string    `json:"init_oldest_date"`
	InitStartDate    string    `json:"init_start_date"`
	UpdatedTime      time.Time `json:"updated_time"`
}

type CloudAccountAtt struct {
	AccountAttId       int    `json:"acccount_att_id"`
	AccountId          int    `json:"account_id"`
	AccountAttKey      string `json:"account_att_key"`
	AccountAttKeyName  string `json:"account_att_key_name"`
	AccountAttValue    string `json:"account_att_value"`
	AccountAttRowIndex int    `json:"account_att_row_index"`
	Deleted            int    `json:"deleted"`
}

type BillAttr struct {
	BillAttrId     int       `json:"bill_att_id"`
	AccountId      int       `json:"account_id"`
	InitOldestDate string    `json:"init_oldest_date"`
	InitStartDate  string    `json:"init_start_date"`
	InitEndDate    string    `json:"init_end_date"`
	Deleted        int       `json:"deleted"`
	CreatedTime    time.Time `json:"created_time"`
	UpdatedTime    time.Time `json:"updated_time"`
}
