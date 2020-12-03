package main

import "fmt"

func foogo() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	ch := make(chan struct{}, 0)

	for _, v := range s {
		go func(v int) {
			fmt.Println(v) // ok, v是闭包的形参，值与循环变量一致
			ch <- struct{}{}
		}(v) // 将循环变量作为实参传入
	}
	for range s {
		<-ch
	}
}

func fooor() {
	a := [3]int{1, 2, 3}
	var resc []func()
	for i, v := range a {
		a[i] = 0
		fmt.Println("for:", a)
		m, n := i, v // 所有在循环体里创建的匿名函数都捕捉了这两个变量本身（变量的地址），而不是捕捉了这两个变量的值。
		resc = append(resc, func() {
			fmt.Println("m, n:", m, n)
			a[m] = n
		})
		fmt.Println("for1:", a)
	}
	fmt.Println("==============:", a)
	for _, rec := range resc {
		fmt.Println("for4:", a)
		rec()
		fmt.Println("for5:", a)
	}
	fmt.Println("==============:", a)
}
