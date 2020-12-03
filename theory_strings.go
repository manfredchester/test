package main

import (
	"fmt"
	"strings"
)

func Str() {
	fmt.Println(strings.IndexFunc("m*+&^]&./", s))
}

func s(c rune) bool {
	// if c != "]" {
	if c != ']' {
		return false
	}
	return true
}
