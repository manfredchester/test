package main

import (
	"fmt"
	"runtime"
	"time"
)

func theoryRountine() {
	maxProcs()
	// rountine()
}

func maxProcs() {
	// 查询当前CPU个数
	fmt.Println(runtime.NumCPU())
	// 设定P值 默认CPU核数，充分使用每一个CPU，提高并发性能
	fmt.Println(runtime.GOMAXPROCS(1))
	// GOMAXPROCS其实并不影响gorountine的数量而是决定了会有多少cpu去服务于这个进程
	ch := make(chan struct{}, 0)
	// 查看当前routine数量
	fmt.Println(runtime.NumGoroutine())

	go func() {
		for {
		}
	}()
	go func() {
		for {
		}

	}()
	go func() {
		for {
		}

	}()
	go func() {
		for {
		}

	}()
	go func() {
		for {
		}

	}()
	fmt.Println(runtime.NumGoroutine())

	<-ch
}

func rountine() {
	var i int
	ch := make(chan int, 0)
	go func() {
		for {
			// if i < 1000 {
			i++
			ch <- 1
			// }
		}
	}()
	go func() {
		for {
			<-ch
			fmt.Println(i)
		}
	}()
	fmt.Println(time.Now().Unix())
	// 两者等价
	// <-time.NewTicker(time.Duration(1 * time.Second)).C
	<-time.After(time.Duration(1 * time.Second))
	fmt.Println(time.Now().Unix())

	wa := make(chan int, 0)
	// <-wa

	se := make(chan struct{}, 0)
	go func() {
		<-time.After(time.Duration(10 * time.Second))
		wa <- 1
	}()
	select {
	case <-se:
		fmt.Println("se")
	case <-wa:
		fmt.Println("wa")
	}
}
