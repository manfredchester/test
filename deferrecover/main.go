package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
	ch := make(chan struct{})
	cx, _ := context.WithTimeout(context.Background(), time.Second)
	// routine 泄露
	go func() {
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
		fmt.Println("goroutine 结束")
	}()

	select {
	case <-ch:
		fmt.Println("res")
	case <-cx.Done():
		fmt.Println("timeout")
	}
	time.Sleep(5 * time.Second)
	// func1
	Go(func() { panic("错误失败") })
	// func2
	Go(func() { panic("请求错误") })
	time.Sleep(10 * time.Second)
}

func Go(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("错误:%v\n", err)
		}
	}()
	fn()
}
