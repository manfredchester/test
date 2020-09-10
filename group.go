package main

import (
	"fmt"
	"os"
)

func group1() {
	// panic occured case1: err deal
	field1, err := field1()
	// panic occured case2: index out of range
	fmt.Println(field1[4], err)
	// panic occured case3: err deal of origin package
	f, err := os.Open("test.txt")
	fmt.Println(f, err)
	// panic occured case4: conversion
	var case3 interface{}
	fmt.Println(case3.(string))
}

func group2() {
	// any other panic errors occured
}

func field1() ([]string, error) {
	field1 := []string{"test"}
	return field1, nil
}

func groupA() {
	// this defer&reovce because of BaseA function's required
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e.(error).Error())
		}
	}()
	// fieldA()
	panic("errors occured")
}

func groupB() {
	// any work
}
