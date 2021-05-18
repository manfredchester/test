package main

import "fmt"

func gooo() {
	goofunc3(f1)
	goofunc3(f2, "ss")
	goofunc3(f3, "ss", "ss")

}

// func goofunc1(f func()) {

// }

// func goofunc2(f func(interface{})) {

// }

func goofunc3(f interface{}, args ...interface{}) {
	if len(args) > 1 {
		go f.(func(...interface{}))(args)
	} else if len(args) == 1 {
		go f.(func(interface{}))(args[0])
	} else {
		go f.(func())()
	}
}

func f1() {
	fmt.Println("f1 done")
}

func f2(i interface{}) {
	fmt.Println("f2 done", i)
}

func f3(args ...interface{}) {
	fmt.Println("f3 done", args)
}

func f4(args []interface{}) {
	fmt.Println("f4 done", args)
}
