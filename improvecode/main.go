package main

import "fmt"

func main() {
	// test1
	// 减小程序计算量
	// 把乘法换成加法，以n为步长，这样就减小了外循环的代码量。
	var n = 3
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for i := 0; i < n; i++ {
		ni := n * i
		for j := 0; j < n; j++ {
			a[ni+j] = b[j]
		}
	}
	fmt.Println(a)
	fmt.Println(b)
	// n*i=n+n+n+.....n
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var ni = 0
	for i := 0; i < n; i++ {
		// ni := n * i
		for j := 0; j < n; j++ {
			a[ni+j] = b[j]
		}
		ni += n
	}
	fmt.Println(a)
	fmt.Println(b)

	// 提取代码中的公共部分
	// 编译后只有一个乘法。减少了6个时钟周期（一个乘法周期大约为3个时钟周期）。
	var x, y, z int
	up := sum((x-1)*z + y)
	down := sum((x+1)*z + y)
	left := sum(x*z + y - 1)
	right := sum(x*z + y + 1)
	_ = up + down + left + right
}

func sum(int) int { return 0 }
