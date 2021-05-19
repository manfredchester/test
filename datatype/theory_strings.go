package main

import (
	"fmt"
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
