package main

import "fmt"

func theoryReturn() {
	fmt.Println(case1())
	fmt.Println(case2())
	fmt.Println(case3())
	fmt.Println(case4())
	// fmt.Println(case1())

}

func case1() (i int) {
	defer func() {
		i++
	}()
	i = 10
	return
}

func case2() (i int) {
	defer func() {
		i++
	}()
	var b *int
	i = 10
	b = &i
	fmt.Println(b)
	return i
}

func case3() (i int) {
	i = 10
	defer func() {
		i++
	}()
	return 2
}

func case4() (i *int) {
	var a int
	defer func() {
		*i++
	}()
	a = 10
	i = &a
	return
}
