package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
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
	// Init()
	// if cmd.UUID == "" || cmd.File == "" || cmd.CldType == "" || cmd.CldID == "" {
	// 	fmt.Println("以下命令行参数必填，且信息均为真实有效")
	// 	flag.PrintDefaults()
	// 	return
	// }
	// fmt.Println(fmt.Sprintf("command: %+v", cmd))
	// AzureBillTrans(cmd)
	mysqlRule7()
}
func mysqlRule7() {
	result := true
	message := ""
	mysqlCnf := "/etc/my.cnf"

	cmd1 := exec.Command("cat", mysqlCnf)
	cmd2 := exec.Command("grep", "skip-grant-tables")
	var outbuf1 bytes.Buffer
	cmd1.Stdout = &outbuf1
	if err := cmd1.Start(); err != nil {
		message += err.Error()
		result = false
		fmt.Println("message1:", message)
		fmt.Println("result1:", result)
		return
	}
	if err := cmd1.Wait(); err != nil {
		message += err.Error()
		result = false
		fmt.Println("message2:", message)
		fmt.Println("result2:", result)
		return
	}
	var outbuf2 bytes.Buffer
	cmd2.Stdin = &outbuf1
	cmd2.Stdout = &outbuf2
	if err := cmd2.Start(); err != nil {
		message += err.Error()
		result = false
		fmt.Println("message3:", message)
		fmt.Println("result3:", result)
		return
	}
	if err := cmd2.Wait(); err != nil {
		message += err.Error()
		result = false
		fmt.Println("message4:", message)
		fmt.Println("result4:", result)
		return
	}
	if outbuf2.String() != "" {
		result = false
	}

	fmt.Println("message5:", message)
	fmt.Println("result5:", result)

	return
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
