# Go语言参考

该部分是对Go语法参考的学习和解读

* 最新版本的Go语言参考: https://go.dev/ref/spec
* 无泛型版本的Go语言参考(1.17): https://go.dev/doc/go1.17_spec

## 清单
- [字符串: string，[]rune和[]byte的区别](./string-and-rune/string_and_rune.go)
- [常量定义: const和iota的使用](./const-and-iota/const_and_iota.go)
- [type关键字: 类型定义，类型别名，内嵌](./type/type.go)
- [struct关键字: 不同结构体的区别](./struct-diff/struct_diff.go)
- [nil关键字: 有类型的nil和无类型的nil](./nil/nil.go)
- [空struct的应用](./empty-struct/README.md) 空结构体`struct{}`的大小为0，且具有相同的内存地址