package main

import (
	"fmt"
	"reflect"
)

type Empty struct{}

var (
	g1 struct{ a int }
	g2 struct{ b int }
	g3 struct{ a int }
)

func main() {
	var (
		e1 Empty
		e2 struct{}
		v1 struct{ a int }

		// struct的成员相同，定义的顺序不同，也是两种不同的类型
		v2 struct {
			a int
			b string
		}
		v3 struct {
			b string
			a int
		}
		v4 struct {
			a int
			b string
		}
	)

	gt1 := reflect.TypeOf(g1)
	gt2 := reflect.TypeOf(g2)
	gt3 := reflect.TypeOf(g3)
	et1 := reflect.TypeOf(e1)
	et2 := reflect.TypeOf(e2)
	t1 := reflect.TypeOf(v1)
	t2 := reflect.TypeOf(v2)
	t3 := reflect.TypeOf(v3)
	t4 := reflect.TypeOf(v4)

	// struct { a int } struct { b int } struct { a int } struct { a int }
	fmt.Println(gt1, gt2, gt3, t1)
	fmt.Println(gt1 == gt2, gt1 == gt3, gt1 == t1) // false true true
	fmt.Println(et1 == et2)                        // false
	fmt.Println(t2 == t3, t2 == t4)                // false true

	// 类型相同，那么就可以互相赋值
	g1.a = 1
	fmt.Println(g1, g3, v1) // {1} {0} {0}
	v1 = g1
	g1.a = 2
	g3 = g1
	fmt.Println(g1, g3, v1) // {2} {2} {1}

}
