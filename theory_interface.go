package main

func TT(vv []interface{}) string {
	return "dsd"
}
func TT2(vv interface{}) string {
	switch vv.(type) {
	case string:
	case []string:
	}
	return "dsd"
}
func slice_interface() {
	ss := []int{1, 2, 3}
	// go 没有帮助我们自动把 slice 转换成 interface{} 类型的 slice，所以出错了。
	// go 不会对 类型是interface{} 的 slice 进行转换 。
	// 1. []interface 代表了 元素为interface 的slice，相当于指明了入参的输入类型
	// 2. []interface 在底层已经设置了存储空间值，interface是2个字节长度的空间，所以固定长度N*2，而入参的其他数据类型的slice是根据是输入类型而判断的,N*sizeof(T)
	tt := make([]interface{}, 0)
	for _, v := range ss {
		tt = append(tt, v)
	}
	TT(tt)

	ss2 := 2

	TT2(ss2)

	TT2(ss)
}
