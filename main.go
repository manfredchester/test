package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"test/zhlog"
	"time"
)

func base1() {
	// this defer&reovce because of main function's required
	defer func() {
		// Capture any errors that occur and do not propagate in the opposite direction of the call stack
		if e := recover(); e != nil {
			fmt.Println(e.(error).Error())
		}
	}()
	// base1 function:
	// I hope to end all work immediately in case of an error at ant time at ant function
	group1()
	group2()
}

func baseA() {
	// this defer&reovce because of main function's required
	defer func() {
		// Capture any errors that occur and do not propagate in the opposite direction of the call stack
		if e := recover(); e != nil {
			fmt.Println(e.(error).Error())
		}
	}()
	// baseA fucntion:
	// I hope to Encounter an error and continue to work at groupA()
	groupA()
	// groupA() error occured but not impacted on next groupB() execution
	// I hope to end all work immediately in case of an error at ant time at groupB()
	groupB()
}

func Test(e Empty) {
	var v []interface{}
	e.Add(v)
}

// func (e *Empty) Test2() {

// }

type Empty interface {
	Add(...interface{})
	Remove(interface{})
}

type Set struct {
	m map[interface{}]struct{}
}

func (s *Set) Add(ele ...interface{}) {
	for _, item := range ele {
		s.m[item] = struct{}{}
	}
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}
func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}
func (s *Set) Size() int {
	return len(s.m)
}
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

// func (s *Set) Sub(other *Set) {
// 	if len()
// }

func New(items ...interface{}) (s *Set) {
	s = &Set{}
	s.m = make(map[interface{}]struct{}, 0)
	s.Add(items)
	return
}

func main() {
	var e Set
	Test(&e)
	// ss := Set{
	// 	dsa: map[interface{}]struct{},
	// }
	// main function: requirment
	// I hope to Never break down at any time
	// base1()
	// baseA()
	// L1()
	// GMP()
	// zhlog.Error("traceID", "test():%v", test())
	zhlog.Trace("[ERROR]", test())
	// f := Foo{
	// 	I: implOfI{},
	// 	J: implOfJ{},
	// }
	// println(f.String())

	// type A interface {
	// 	a()
	// 	String() string
	// }
	// type B interface {
	// 	b()
	// 	String() string
	// }
	// //  ----- (1)go1.14
	// type O interface {
	// 	A
	// 	B
	// }
	// // ---- (2)go1.13
	// type O2 interface {
	// 	a()
	// 	b()
	// 	String() string
	// }
	// zheng("test")
	// \w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}

	// reg := regexp.MustCompile(`^([A-Za-z]|[\u4e00-\u9fa5])+(\d|_|-|[A-Za-z]|[\u4e00-\u9fa5])+`)
	// fmt.Println(reg.MatchString("QWE"))

	// fmt.Println(Multiple3And5(10))
	// Init()
	// if cmd.UUID == "" || cmd.File == "" || cmd.CldType == "" || cmd.CldID == "" {
	// 	fmt.Println("以下命令行参数必填，且信息均为真实有效")
	// 	flag.PrintDefaults()
	// 	return
	// }
	// fmt.Println(fmt.Sprintf("command: %+v", cmd))
	// AzureBillTrans(cmd)
	// mysqlRule7()
}
func L1() {
	defer func() {
		fmt.Println("L1 defer func invoked")
	}()
	fmt.Println("L1  invoked")
	L2()
	fmt.Println("do something after L2 in L1")
}

func L2() {
	defer func() {
		fmt.Println("L2 defer func invoked")
		if e := recover(); e != nil {
			fmt.Println("i get you")
		}
	}()
	fmt.Println("L2  invoked")
	L3()
	fmt.Println("do something after L3 in L2")
	panic("runtime exception")
}

func L3() {}
func GMP() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	println("Connext PanCloud NO.1!")
}

type I interface {
	f()
	String() string
}

type implOfI struct{}

func (implOfI) f() {}
func (implOfI) String() string {
	return "implOfI"
}

type J interface {
	g()
	String() string
}

type implOfJ struct{}

func (implOfJ) g() {}
func (implOfJ) String() string {
	return "implOfJ"
}

type Foo struct {
	I
	J
}

func (Foo) String() string {
	return "Foo"
}

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
