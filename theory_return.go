package main

import "fmt"

func theoryReturn() {
	fmt.Println(case0())  //10
	fmt.Println(case1())  //11
	fmt.Println(case2())  //10 11
	fmt.Println(case3())  //3
	fmt.Println(*case4()) //11
}

func case0() int {
	i := 10
	defer func() {
		i++
	}()
	return i
}

// 匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值
// ‍defer、return、返回值三者的执行顺序应该是：return最先给返回值赋值；接着defer开始执行一些收尾工作；最后RET指令携带返回值退出函数。
func case1() (i int) {
	defer func() {
		i++
	}()
	i = 10
	return
}

func case2() (i int) {
	defer func() {
		i++
	}()
	var b *int
	i = 10
	b = &i
	// fmt.Println(b)
	fmt.Print(*b)
	return i
}

func case3() (i int) {
	i = 10
	defer func() {
		i++
	}()
	return 2
}

func case4() (i *int) {
	var a int
	defer func() {
		*i++
	}()
	a = 10
	i = &a
	return
}
