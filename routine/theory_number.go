package main

import "fmt"

// 第二个 goroutine 运行到 i = 10 的时候，第一个 goroutine 中的 for 循环已经结束退出了，因此 A 通道再也没有了接收者，而第二个 goroutine 里还继续向 A 通道写数据，当然就出错了。
func main() {
	A := make(chan bool)
	B := make(chan bool)
	Exit := make(chan bool)
	go func() {
		for i := 1; i <= 4; i += 2 {
			if ok := <-A; ok {
				fmt.Println("A:", i)
				B <- true
			}
		}
		fmt.Println("=====byebye===A==")
	}()
	go func() {
		defer func() {
			Exit <- true
		}()
		for i := 2; i <= 4; i += 2 {
			if ok := <-B; ok {
				fmt.Println("B:", i)
				// if i != 4 {
				A <- true
				// }
			}
		}
		fmt.Println("=====byebye===B==")
	}()
	A <- true
	<-Exit
}
