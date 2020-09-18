package main

import (
	"fmt"
	"runtime"
)

func proc() {
	runtime.GOMAXPROCS(1)
	var i int
	// go func() {
	// 	i++
	// }()
	go func() {
		fmt.Println(i)
	}()
	// <-time.NewTicker(time.Duration(3600))
	// select {}
	// time.Sleep(3600)
	ch := make(chan struct{}, 0)
	<-ch
}
