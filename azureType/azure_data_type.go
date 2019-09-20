package azureType

import (
	"time"

	"pub/request"
)

type InitDataTimeReq struct {
	AccountId     int    `json:"account_id"`
	InitStartTime string `json:"init_start_time"`
}

type BillSyncDetails struct {
	BillSyncDetailsId int       `json:"bill_sync_details_id"` //账单同步详情ID
	AccountId         int       `json:"account_id"`           //账号ID
	SyncDate          string    `json:"sync_date"`            //同步年月
	SyncState         int       `json:"sync_state"`           //同步状态
	SyncRetry         int       `json:"sync_retry"`           //重试次数
	Deleted           int       `json:"deleted"`              //删除标志
	CreatedBy         string    `json:"created_by"`           //创建人
	CreatedTime       time.Time `json:"created_time"`         //创建时间
	UpdatedBy         string    `json:"updated_by"`           //更新人
	UpdatedTime       time.Time `json:"updated_time"`         //更新时间
}

type BillCppReq struct {
	CppId        int                        `json:"cpp_id"`
	AccountId    int                        `json:"account_id"`
	CppName      string                     `json:"cpp_name"`
	ServiceType  string                     `json:"service_type"`
	StartDate    string                     `json:"start_date"`
	EndDate      string                     `json:"end_date"`
	LicenceCount int                        `json:"licence_count"`
	MonthCost    float64                    `json:"month_cost"`
	Pagination   *request.RequestPagination `json:"pagination"`
}

// cpp订单信息 bill_cpp
type BillCpp struct {
	CppId     int    `json:"cpp_id"`     //CPPID
	AccountId int    `json:"account_id"` //账号ID
	CppName   string `json:"cpp_name"`   //CppName

	ServiceType  string  `json:"service_type"`  //ServiceType
	StartDate    string  `json:"start_date"`    //StartDate
	EndDate      string  `json:"end_date"`      //EndDate
	LicenceCount int     `json:"licence_count"` //LicenceCount
	MonthCost    float64 `json:"month_cost"`    //MonthCost

	Deleted     int       `json:"deleted"`      //删除标志
	CreatedBy   string    `json:"created_by"`   //创建人
	CreatedTime time.Time `json:"created_time"` //创建时间
	UpdatedBy   string    `json:"updated_by"`   //更新人
	UpdatedTime time.Time `json:"updated_time"` //更新时间
}

type BillAzureCppReqAll struct {
	BillAccountUuid string            `json:"bill_account_uuid"`
	CppList         []BillAzureCppReq `json:"cpp_list"`
}

type BillAzureCppReq struct {
	Uuid            string                     `json:"uuid"`
	BillAccountUuid string                     `json:"bill_account_uuid"`
	AccountName     string                     `json:"account_name"`
	ServiceType     string                     `json:"service_type"`
	StartDate       string                     `json:"start_date"`
	EndDate         string                     `json:"end_date"`
	LicenceCount    int                        `json:"licence_count"`
	MonthCost       float64                    `json:"month_cost"`
	Pagination      *request.RequestPagination `json:"pagination"`
}

type BillAzureCpp struct {
	Uuid            string    `json:"uuid"`
	BillAccountUuid string    `json:"bill_account_uuid"`
	AccountName     string    `json:"account_name"`
	ServiceType     string    `json:"service_type"`
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
	LicenceCount    int       `json:"licence_count"`
	MonthCost       float64   `json:"month_cost"`
	Deleted         int       `json:"deleted"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

type BillCloudMaster struct {
	ItemUuid string

	BillAccountUuid string
	CloudType       string
	AccountId       string
	BillingCycle    string

	Date             string
	SubscriptionGuid string
	SubscriptionName string
	Product          string
	MeterCategory    string
	ConsumedQuantity string
	ExtendedCost     string
	Tags             string
	ResourceGroup    string

	Deleted    int
	CreateTime time.Time
	UpdateTime time.Time
}

type BillCloudTag struct {
	Uuid            string
	ItemUuid        string
	BillAccountUuid string
	BillingCycle    string
	AccountId       string
	TagKey          string
	TagValue        string
	TagType         string

	Deleted    int
	CreateTime time.Time
	UpdateTime time.Time
}
