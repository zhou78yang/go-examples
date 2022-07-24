package main

import "fmt"

type MustKeyStruct struct {
	Name string
	Age  int
	_    struct{}
}

func main() {
	// 只能使用key的方式进行初始化
	a := MustKeyStruct{
		Name: "zhangsan", Age: 18,
	}
	//a = MustKeyStruct{"lisi", 22} // 编译失败
	fmt.Println(a)
}
