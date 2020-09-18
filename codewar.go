package main

import "fmt"

func S3(number int) (res int) {
	n := (number - 1) / 3
	fmt.Println(n)

	res = n*1 + n*(n-1)/2*3
	fmt.Println(res)

	return
}
func S5(number int) (res int) {
	n := (number - 1) / 5
	fmt.Println(n)

	res = n*1 + n*(n-1)/2*5
	fmt.Println(res)

	return
}
func S15(number int) (res int) {
	n := (number - 1) / 15
	fmt.Println(n)

	res = n*1 + n*(n-1)/2*15
	fmt.Println(res)
	return
}
func Multiple3And5(number int) int {
	// fmt.Println(10 / 3)
	return S3(number) + S5(number) - S15(number)
	// return 0
}
func test() string {
	tettt()
	return "sad"
}
func tettt() {
	panic("sad1")
}
