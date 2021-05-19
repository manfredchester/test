package main

import "fmt"

func Arr() {
	fmt.Println("=========数组与切片===========")
	arr := [2]int{1}
	// 固定长度
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	// 默认值
	fmt.Println(arr)

	sli := []int{1}
	fmt.Println(len(sli))
	fmt.Println(cap(sli))
	fmt.Println(sli)
	sli = append(sli, 2, 3)
	// 自增
	fmt.Println(len(sli))
	// cap为前一数值的2倍
	fmt.Println(cap(sli))
	fmt.Println(sli)
	fmt.Println("=========切片为添加元素===========")
	sli2 := []int{1, 2, 3}
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	fmt.Println(sli2)
	// 切片为添加元素
	sli2 = append(sli2, []int{4, 5}...)
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	fmt.Println(sli2)
	// cap为前一数值的2倍
	sli2 = append(sli2, []int{6, 7}...)
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	fmt.Println(sli2)

	fmt.Println("=========copy slice===========")
	nnslice := make([]int, len(sli2), cap(sli2))
	fmt.Println(copy(nnslice, sli2))
	fmt.Println("====================", nnslice)
	nnslice[0] = -100
	fmt.Println("====================", nnslice)
	nnslice = append(nnslice, -200)
	fmt.Println("====================", nnslice)

	fmt.Println("=========newslice===========")
	newslice := sli2[1:3]
	fmt.Println("====================", newslice)
	fmt.Println(sli2)
	newslice[0] = -1
	fmt.Println("====================", newslice)
	fmt.Println(sli2)
	newslice = append(newslice, -2)
	fmt.Println("====================", newslice)
	fmt.Println(sli2)
}

// 使用切片操作符切取切片时，上界是切片的容量，而非长度。
func cap1() {
	array := [10]uint32{1, 2, 3, 4, 5}
	s1 := array[:5]

	s2 := s1[5:10]

	fmt.Println(s1)
	fmt.Println(s2)

	s1 = append(s1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}
