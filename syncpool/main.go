package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type Person struct {
		Name string
	}
	pool := sync.Pool{
		New: func() interface{} {
			return new(Person)
		},
	}
	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)
	p2 := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p2)

	p3 := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p3)
	// 在实际的并发使用场景中，无法保证这种顺序，最好的做法是在 Put 前，将对象清空
	p.Name = "first Name"
	pool.Put(p)
	time.Sleep(2 * time.Second)

	p2.Name = "second Name"
	pool.Put(p2)
	time.Sleep(2 * time.Second)
	p3.Name = "third Name"
	pool.Put(p3)

	fmt.Println("1 调用 Get: ", pool.Get().(*Person))
	fmt.Println("2 调用 Get: ", pool.Get().(*Person))
	fmt.Println("3 调用 Get: ", pool.Get().(*Person))

}
