package zhlog

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var DEBUG_TARGETS []string
var rv *regexp.Regexp

// 日志表 s_log
type SLog struct {
	LogId         int       `json:"log_id"`          //日志ID
	LogResId      int       `json:"log_res_id"`      //资源ID
	LogResName    string    `json:"log_res_name"`    //资源名称
	LogTraingId   string    `json:"log_traing_id"`   //业务跟踪ID
	LogBusinessId string    `json:"log_business_id"` //业务主键
	LogMicroUri   string    `json:"log_micro_uri"`   //微服务URI
	LogLevel      string    `json:"log_level"`       //日志级别
	LogType       string    `json:"log_type"`        //日志类型
	LogMessage    string    `json:"log_message"`     //日志内容
	LogIp         string    `json:"log_ip"`          //请求IP
	LogTakeTime   int       `json:"log_take_time"`   //耗时
	AppId         int       `json:"app_id"`          //应用ID
	OrgId         int       `json:"org_id"`          //所属机构
	Deleted       int       `json:"deleted"`         //删除标志
	CreatedBy     string    `json:"created_by"`      //创建人
	CreatedTime   time.Time `json:"created_time"`    //创建时间
}

func init() {
	rv = regexp.MustCompile(`.func\d+(.\d+)?\s*$`)
}

func UUID(n int) string {
	const charMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		ch := buf[i]
		buf[i] = charMap[int(ch)%62]
	}
	return string(buf)

}

// func Error(err error) {
// 	fmt.Fprintln(os.Stderr, Trace(err.Error()))
// }

// func Log(msg string, args ...interface{}) {
// 	msg = strings.TrimRightFunc(fmt.Sprintf(msg, args...), unicode.IsSpace)
// 	fmt.Println(msg)
// }

//Error output message with stack trace. prefix is added to every single
//line of log output for tracing purpose.
func Error(traceID, msg string, args ...interface{}) {
	log(traceID, Trace("[ERROR]"+msg, args...).Error())
}

func ErrorToDBPro(traingID, businessID, logType string, logTakeTime int, msg string, args ...interface{}) {
	message := Trace("[ERROR]"+msg, args...).Error()
	if len(traingID) > 0 {
		message += "\nTRACE_ID:" + traingID
	}
	session := getSession()
	_, err := session.InsertOne(&SLog{
		// LogResId      :,
		// LogResName    :,
		LogTraingId:   traingID,
		LogBusinessId: businessID,
		LogMicroUri:   "api-service-platform",
		LogLevel:      "ERROR",
		LogType:       logType,
		LogMessage:    message,
		// LogIp         :,
		LogTakeTime: logTakeTime,
		// AppId         :,
		// OrgId         :,
		CreatedTime: time.Now(),
	})
	if err != nil {
		session.Rollback()
		panic(err)
	}
	session.Commit()
}

func ErrorToDB(traingID, businessID, logType string, logTakeTime int, msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	session := getSession()
	_, err := session.InsertOne(&SLog{
		// LogResId      :,
		// LogResName    :,
		LogTraingId:   traingID,
		LogBusinessId: businessID,
		LogMicroUri:   "api-service-platform",
		LogLevel:      "ERROR",
		LogType:       logType,
		LogMessage:    msg,
		// LogIp         :,
		LogTakeTime: logTakeTime,
		// AppId         :,
		// OrgId         :,
		CreatedTime: time.Now(),
	})
	if err != nil {
		session.Rollback()
		panic(err)
	}
	session.Commit()
}

func InfoToDB(traingID, businessID, logType string, LogTakeTime int, msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	session := getSession()
	_, err := session.InsertOne(&SLog{
		// LogResId      :,
		// LogResName    :,
		LogTraingId:   traingID,
		LogBusinessId: businessID,
		LogMicroUri:   "api-service-platform",
		LogLevel:      "INFO",
		LogType:       logType,
		LogMessage:    msg,
		// LogIp         :,
		LogTakeTime: LogTakeTime,
		// AppId         :,
		// OrgId         :,
		CreatedTime: time.Now(),
	})
	if err != nil {
		session.Rollback()
		panic(err)
	}
	session.Commit()
}

//Log output message. prefix is added to every single line of log
//output for tracing purpose.
func Log(traceID, msg string, args ...interface{}) {
	log(traceID, "[INFO]"+msg, args...)
}

func Dbg(msg string, args ...interface{}) {
	if len(DEBUG_TARGETS) == 0 {
		return
	}
	var wanted bool
	caller := ""
	log := Trace("")
	for _, l := range log {
		if l != "" {
			caller = l
			break
		}
	}
	caller = rv.ReplaceAllString(caller, "")
	if DEBUG_TARGETS[0] == "*" {
		wanted = true
	} else {
		if caller == "" {
			wanted = true
		} else {
			for _, t := range DEBUG_TARGETS {
				if strings.HasSuffix(caller, t) {
					wanted = true
					break
				}
			}
		}
	}
	if wanted {
		Log("Dbg", strings.TrimSpace(caller)+"> "+msg, args...)
	}
}

func SetDebugTargets(targets string) {
	DEBUG_TARGETS = []string{}
	for _, t := range strings.Split(targets, ",") {
		t = strings.TrimSpace(t)
		if t != "" {
			DEBUG_TARGETS = append(DEBUG_TARGETS, t)
		}
	}
}

func Perf(tag string, work func()) {
	start := time.Now()
	Dbg("[EXEC]%s", tag)
	work()
	elapsed := time.Since(start).Seconds()
	Dbg("[DONE]%s (elapsed: %f)", tag, elapsed)
}
