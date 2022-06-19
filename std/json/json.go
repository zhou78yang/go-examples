package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func blank(funcName string) {
	fmt.Println(strings.Repeat("=", 20), funcName, strings.Repeat("=", 20))
}

func main() {
	blank("marshal")
	marshal()

	blank("unmarshal")
	unmarshal()

	blank("decodeInterface")
	decodeInterface()

	blank("jsonTag")
	jsonTag()
}

// marshal 通过`json.Marshal`方法序列化一个Go对象为json数据（[]byte）类型
func marshal() {
	msg := struct {
		Name string
		Body string
		Time int64
	}{
		"zhangsan", "Hello", 123456789,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	if string(b) == `{"Name":"zhangsan","Body":"Hello","Time":123456789}` {
		fmt.Println(string(b))
	}
}

// unmarshal 通过`json.Unmarshal`对JSON数据进行反序列化
func unmarshal() {
	var msg struct {
		Name string
		Body string
		Time int64
	}

	b := []byte(`{"Name":"zhangsan","Body":"Hello","Time":123456789}`)
	err := json.Unmarshal(b, &msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}

// decodeInterface 结构不确定的JSON数据反序列化
func decodeInterface() {
	b := []byte(`{"name": "zhangsan", "age": 18, "weight": 62.1, "tags": ["male", "student"], "is_adult": true,
				  "addr": {"country": "cn", "city": "sz"}, "nil_value": null}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	fmt.Println("object:", f)

	m, ok := f.(map[string]interface{})
	if !ok {
		panic("[genericJson] decode map error")
	}

	for k, v := range m {
		switch v := v.(type) {
		case bool:
			fmt.Println(k, "is bool:", v)
		case string:
			fmt.Println(k, "is string:", v)
		case float64:
			fmt.Println(k, "is number:", v)
		case []interface{}:
			fmt.Println(k, "is array:", v)
		case map[string]interface{}:
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println(k, "marshal err:", err)
			}
			fmt.Println(k, "is object:", string(b))
		case nil:
			fmt.Println(k, "is null")
		default:
			fmt.Printf("%s is unknown type %T\n", k, v)
		}
	}
}

// jsonTag struct通过json tag
func jsonTag() {
	type User struct {
		Id       int    `json:"-"` // 不参与序列化和反序列化
		Name     string `json:"name"`
		Age      int    `json:"age,string"` // json字段名age，json总类型为string
		Address  string // 无json tag也会被序列化成键名为Address的键值对
		Desc     string `json:"desc,omitempty"`      // 零值时不序列化
		EmptyInt int    `json:"empty_int,omitempty"` // 注意，会导致0不序列化
		Neg      string `json:"-,"`                  // 用来序列化key为"-"的情况
	}

	users := []User{
		{Id: 1, Name: "张三", Age: 18},
		{Id: 2, Name: "李四", Age: 22, EmptyInt: 11, Desc: "lisi is worker"},
	}

	b, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
