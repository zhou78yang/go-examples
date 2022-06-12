package main

import (
	"fmt"
	"reflect"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "my error"
}

func getError(flag bool) error {
	var err *MyError = nil
	if flag {
		err = &MyError{}
	}
	return err
}

func getError2(flag bool) error {
	if flag {
		return &MyError{}
	}
	return nil
}

func main() {
	// interface的值由两部分组成（类型和数据），只有这两者同时为nil才表示接口是nil
	// 反例: 下面返回带类型的nil，if条件永远为true
	err := getError(false)
	if err != nil {
		fmt.Println("getError", reflect.TypeOf(err))
	}

	// 正例: 返回无类型的nil
	err = getError2(false)
	if err != nil {
		fmt.Println("getError2", reflect.TypeOf(err))
	}

}
