package main

import (
	"fmt"
)

func main() {
	// 常量在定义时必须赋值，不指定类型时可以自动推导类型
	const pi = 3.1415
	var f1 float32 = pi
	var f2 float64 = pi
	fmt.Printf("%f %T %T\n", pi, f1, f2) // 3.1415 float32 float64

	// const声明多个常量，如果省略值，就表示和上一行值相同(包括格式)
	const (
		c1 int = 100
		c2
		c3 = "aaa"
		c4
		c5, c6 = iota, iota + 1 // 此时iota=4
		_, _                    // 此时iota=5
		c7, _                   // 此时iota=6
	)
	fmt.Println(c1, c2, c3, c4) // 100 100 aaa aaa
	fmt.Println(c5, c6, c7)     // 4 5 6

	// iota是Go语言中的常量计数器，每一个const关键字出现iota都会被重置为0
	const (
		Sun = iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
	)
	fmt.Println(Sun, Mon, Tue, Wed, Thu, Fri, Sat) // 0 1 2 3 4 5 6

	// iota从1开始的方式
	const (
		i1 = iota + 1
		i2
		i3
	)
	fmt.Println(i1, i2, i3) // 1 2 3

	// 移位常量
	const (
		b1 = 1 << iota
		b2
		b3
		b4
		b5
	)
	fmt.Println(b1, b2, b3, b4, b5) // 1 2 4 8 16
}
