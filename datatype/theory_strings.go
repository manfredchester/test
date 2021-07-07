package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Str() {
	fmt.Println(strings.IndexFunc("m*+&^]&./", s))

	str := "Golang梦工厂"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))

	var ans float64 = 15 + 25 + 5.2
	fmt.Println(ans)
}

func s(c rune) bool {
	// if c != "]" {
	if c != ']' {
		return false
	}
	return true
}

func main() {
	jsonStr := `{"name":"tom","user_id":"999","age":"23"}`
	var str string
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(err)

	v := reflect.ValueOf(m)
	keys := v.MapKeys()
	for _, k := range keys {
		v1 := v.MapIndex(k).Interface().(string)
		str += v1
		fmt.Println(str)
	}
	// r := gjson.Parse(jsonStr)
	// r.ForEach(func(key, value gjson.Result) bool {
	// 	fmt.Println(key, value)
	// 	return true
	// })
}
