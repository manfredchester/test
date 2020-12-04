package main

import "fmt"

func gotoloop() {
	for i := 0; i < 10; i++ {
		if i > 3 {
			goto LAbEL2
		}
		fmt.Println("i:", i)
	LAbEL2:
		fmt.Println("LastLAbEL:", i)
	}
}
