package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

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

// rule1
// 以大小字母或中文开头，可包含数字、下划线（_）和连字符（-）。

// rule2
// 长度为2-128个字符，不能以特殊字符及数字开头，只可包含特殊字符中的"."、"_"、"-"和":"

// rule3
// 1）8 - 30 个字符
// 2）必须同时包含三项（大写字母、小写字母、数字、 ()`~!@#$%^&*_-+=|{}[]:;'<>,.?/ 中的特殊符号）
// 3）不能以斜线号（/）为首字符

// rule4
// 由小写字母、数字、下划线组成、字母开头，字母或数字结尾，最长16个字符。
// rule5
// 大写、小写、数字、特殊字符占三种，长度8-32位；特殊字符为！@#￥%……&*()_+-=
// rule6
// 大写、小写、数字、特殊字符占三种，长度8-128位；特殊字符为！@#￥%……&*()_+-=
func zheng(str string) {
	// matched, err := regexp.MatchString(`^([A-Za-z]|[\\u4e00-\\u9fa5])+(\d|_|-|[A-Za-z]|[\\u4e00-\\u9fa5])+`, str)
	// fmt.Println(matched, err)
	// fmt.Println("==========================================================")
	// str := "Z"
	// rule 1
	// var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5a-zA-Z]+(\\d|_|-|[A-Za-z]|[\u4e00-\u9fa5])*$")
	// fmt.Println(hzRegexp.MatchString(str))
	// rule 2
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5a-zA-Z]+(\\d|\\:|\\.|_|-|[A-Za-z]|[\u4e00-\u9fa5]){2,128}$")
	fmt.Println(hzRegexp.MatchString(str))
}

func rule3(str string) bool {
	reVal := false
	startMatch := "^[^/].{7,30}"
	upperMatch := "[A-Z]"
	lowerMatch := "[a-z]"
	numMatch := "[0-9]"
	specialMath := "(\\(|\\)|~|!)"
	reg := regexp.MustCompile(startMatch)
	conditionCount := 0
	if reg.MatchString(str) {
		upperArrary := regexp.MustCompile(upperMatch).FindAllString(str, -1)
		if len(upperArrary) > 0 {
			conditionCount++
		}
		lowerArray := regexp.MustCompile(lowerMatch).FindAllString(str, -1)
		if len(lowerArray) > 0 {
			conditionCount++
		}
		numArray := regexp.MustCompile(numMatch).FindAllString(str, -1)
		if len(numArray) > 0 {
			conditionCount++
		}
		specialArray := regexp.MustCompile(specialMath).FindAllString(str, -1)
		if len(specialArray) > 0 {
			conditionCount++
		}
	}
	if conditionCount >= 3 {
		reVal = true
	}
	return reVal
}
