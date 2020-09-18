package main

import "fmt"

type InD interface {
	InA
	InB
}

type InA interface {
	Int() string
}

type InB interface {
	Double() string
}

func interIn() {
	sta := StA{
		val: "sta",
	}
	fmt.Println(sta.Int())
	s := StA{}
	f(&s)
	std := StD{
		StA: StA{
			val: "int",
		},
		StB: StB{val: "double"},
		My:  "3",
	}
	f2(&std)
}

func f(I InA) {
	fmt.Println("f(I)")
}

func f2(I InD) {
	fmt.Println("f(D)")
}

// struct
type StD struct {
	StA
	StB
	My string
}

func (e *StD) Int() string {
	return "StD Int"
}

type StA struct {
	val string
}

func (e *StA) Int() string {
	return e.val
}

type StB struct {
	val string
}

func (e *StB) Double() string {
	return e.val
}

func (e *StB) Int() string {
	return "stb int()"
}

func interSt() {
	std := StD{
		StA: StA{
			val: "int",
		},
		StB: StB{val: "double"},
		My:  "3",
	}
	fmt.Println(std.Int())
	fmt.Println(std.Double())
	fmt.Println(std.StA.Int())
	fmt.Println(std.StB.Double())
}
