package azureType

import "time"

type UsageReports struct {
	AvailableMonths []ReportMonth
	ContractVersion string
	ObjectType      string
}

type ReportMonth struct {
	Month                           string
	LinkToDownloadSummaryReport     string
	LinkToDownloadDetailReport      string
	LinkToDownloadStoreChargeReport string
	LinkToDownloadPriceSheetReport  string
}

type AzureDetailReport struct {
	AccountOwnerId         string
	AccountName            string
	ServiceAdministratorId string
	SubscriptionId         string
	SubscriptionGuid       string
	SubscriptionName       string
	Date                   string
	Month                  string
	Day                    string
	Year                   string
	Product                string
	MeterId                string
	MeterCategory          string
	MeterSubCategory       string
	MeterRegion            string
	MeterName              string
	ConsumedQuantity       string
	ResourceRate           string
	ExtendedCost           string
	ResourceLocation       string
	ConsumedService        string
	InstanceID             string
	ServiceInfo1           string
	ServiceInfo2           string
	AdditionalInfo         string
	Tags                   string
	StoreServiceIdentifier string
	DepartmentName         string
	CostCenter             string
	UnitOfMeasure          string
	ResourceGroup          string
}

type EnumAzureDetailReport uint8

const (
	AccountOwnerId         = 0
	AccountName            = 1
	ServiceAdministratorId = 2
	SubscriptionId         = 3
	SubscriptionGuid       = 4
	SubscriptionName       = 5
	Date                   = 6
	Month                  = 7
	Day                    = 8
	Year                   = 9
	Product                = 10
	MeterId                = 11
	MeterCategory          = 12
	MeterSubCategory       = 13
	MeterRegion            = 14
	MeterName              = 15
	ConsumedQuantity       = 16
	ResourceRate           = 17
	ExtendedCost           = 18
	ResourceLocation       = 19
	ConsumedService        = 20
	InstanceId             = 21
	ServiceInfo1           = 22
	ServiceInfo2           = 23
	AdditionalInfo         = 24
	Tags                   = 25
	StoreServiceIdentifier = 26
	DepartmentName         = 27
	CostCenter             = 28
	UnitOfMeasure          = 29
	ResourceGroup          = 30
)

type BillAzureDetailReportMetadata struct {
	Uuid            string
	BillAccountUuid string
	BillingCycle    string

	AccountOwnerId         string
	AccountName            string
	ServiceAdministratorId string
	SubscriptionId         string
	SubscriptionGuid       string
	SubscriptionName       string
	Date                   string
	Month                  string
	Day                    string
	Year                   string
	Product                string
	MeterId                string
	MeterCategory          string
	MeterSubCategory       string
	MeterRegion            string
	MeterName              string
	ConsumedQuantity       string
	ResourceRate           string
	ExtendedCost           string
	ResourceLocation       string
	ConsumedService        string
	InstanceId             string
	ServiceInfo1           string
	ServiceInfo2           string
	AdditionalInfo         string
	Tags                   string
	StoreServiceIdentifier string
	DepartmentName         string
	CostCenter             string
	UnitOfMeasure          string
	ResourceGroup          string

	Deleted    int
	CreateTime time.Time
	UpdateTime time.Time
}
