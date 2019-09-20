package main

import (
	"flag"
	"fmt"
	"test/zhlog"
)

// Command Command element
type Command struct {
	UUID    string
	File    string
	CldType string
	CldID   string
}

var cmd Command

// Init init command
func Init() {
	flag.StringVar(&cmd.UUID, "UUID", "", "accountID  such as 88")
	flag.StringVar(&cmd.File, "file", "", "filename  such as process.csv")
	flag.StringVar(&cmd.CldType, "type", "", "cloud type such as azure")
	flag.StringVar(&cmd.CldID, "cldID", "", "cloudid such as V570")
	flag.Parse()
}

func main() {
	Init()
	if cmd.UUID == "" || cmd.File == "" || cmd.CldType == "" || cmd.CldID == "" {
		fmt.Println("以下命令行参数必填，且信息均为真实有效")
		flag.PrintDefaults()
		return
	}
	fmt.Println(fmt.Sprintf("command: %+v", cmd))
	AzureBillTrans(cmd)
}

func AzureBillTrans(cmd Command) {
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error("AzureBillTrans", "%s", e.(error).Error())
		}
	}()
	step1 := InitCsvToDbByMonth(cmd.UUID, cmd.File)
	if !step1 {
		step2 := TransMasterByMonth(cmd.UUID, cmd.CldType, cmd.CldID)
		if !step2 {
			TransTagByMonth(cmd.UUID)
		}
	}

}
