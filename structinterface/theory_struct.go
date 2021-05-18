package main

import "fmt"

type MyIntf interface {
	// Add(...interface{})
	// Remove(interface{})
	Int() int
}

type MyStruct struct {
	val int
}

func (e *MyStruct) Int() int {
	return e.val
}

//定义一个函数，参数为一个自定义的数据类型
func printTheValueByStruct(arg MyStruct) {
	fmt.Printf("the value is %d \n", arg.Int())
}

//定义一个函数，参数为一个接口
func printTheValue(arg MyIntf) int {
	fmt.Printf("the value is %d \n", arg.Int() /*使用接口调用接口的方法*/)
	return arg.Int()
}

func printAnyValue(args ...interface{}) {
	//使用for range方法获取每一个接口
	for _, arg := range args {
		//使用.(type)方法查询接口的数据类型
		switch arg.(type) {
		case int:
			fmt.Println("the type is int")
		case string:
			fmt.Println("the type is string")
		case MyStruct: /*是自定义数据类型*/
			//使用.(数据类型)进行强转
			var b MyStruct = arg.(MyStruct)
			fmt.Println("the type is MyStruct, the function value is ", b.Int() /*d调用数据类型的方法，golang会转换为数据指针类型调用*/, "the struct value is ", b.val /*调用数据结构的数据*/)
		case MyIntf: /*是定义的接口数据类型*/
			fmt.Println("the type is MyIntf interface, the function value is ", arg.(MyIntf).Int() /*将接口强转到指定接口，并调用方法*/)
		}
	}
}

func theoryStruct() {
	var St1 *MyStruct = new(MyStruct)      //创建一个对象指针
	St1.val = 1111111                      //给对象赋值
	var St MyStruct = MyStruct{222222222}  //创建一个对象，给对象赋值
	var a interface{} = MyStruct{33333}    //创建一个对象，将对象赋值后传给一个万能类型接口
	var a1 interface{} = &MyStruct{444444} //创建一个对象，将对象指针传给一个万能类型接口
	fmt.Println("hello world!")

	printTheValue(a1.(MyIntf))          //万能接口a1中放置的对象指针被强制转为MyIntf接口调用
	printTheValueByStruct(a.(MyStruct)) //强制将万能接口a中放置的对象转换为对象传入函数，因为参数是对象

	// 入参接口是指针；入参结构是对象
	printTheValue(St1)          //St1会转换为MyIntf接口被调用其中的方法
	printTheValueByStruct(*St1) //强制将St1的对象使用*显示传入函数，因为参数是对象
	printTheValue(&St)          //将对象的指针传入函数，golang将其转换为Stryc接口
	printTheValueByStruct(St)   //

	printAnyValue(St1, /*传入一个指针，会同MyIntf接口数据类型匹配*/
		St,   /*传入一个对象，会同MyStruct数据类型匹配*/
		*St1, /*将指针显示为对象传入，会同MyStruct数据类型匹配*/
		&St /*将对象的指针传入，会同MyIntf接口数据类型匹配*/)
}
