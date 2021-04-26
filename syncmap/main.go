package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	sm.Store(1, "a")
	sm.Store("2", "b")

	if v, ok := sm.Load(1); ok {
		fmt.Println(v)
	}

	if v, ok := sm.LoadOrStore(1, "A"); ok {
		fmt.Println(v)
	}

	if v, ok := sm.LoadOrStore(33, "CC"); !ok {
		fmt.Println(v)
	}
	sm.Range(func(k, v interface{}) bool {
		fmt.Println(k)
		fmt.Println(v)
		return k != 1
	})
}
