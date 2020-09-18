package main

import (
	"runtime"
	"time"
)

func GMP() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	println("Connext PanCloud NO.1!")
}
