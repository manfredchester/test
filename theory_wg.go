package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func wg() {
	aa := 0
	wg := new(sync.WaitGroup)
	fmt.Println(" main &&&&&", runtime.NumGoroutine())

	wg.Add(1)
	go aaa(wg, aa)
	wg.Wait()
	fmt.Println("finally:", aa)
}

func aaa(wg *sync.WaitGroup, aa int) {
	fmt.Println(" top routine &&&&&", runtime.NumGoroutine())
	defer func() {
		wg.Done()
	}()
	lop := 0
	for {
		if lop == 3 {
			break
		}
		fmt.Println("lop:", lop)
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer func() {
				wg.Done()
			}()
			fmt.Println("&&&&&", runtime.NumGoroutine())
			for i := 0; i < 3; i++ {
				wg.Add(1)
				fmt.Println(" son &&&&&", runtime.NumGoroutine())
				go func(wg *sync.WaitGroup) {
					defer func() {
						wg.Done()
					}()
					fmt.Println("1")
					time.Sleep(time.Second * 3)
					aa++
				}(wg)
			}
		}(wg)
		lop++
	}
}
