package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var object interface{}
	jsonStr := `{"name": "张三", "age": 18, "tags": ["male", "smart", "fat"],
				"friends": [{"name": "lisi", "age": 22}, {"name": "wangwu", "age": 15}],
				"addr": {"country": "CN", "state": "GD", "city": "SZ"}
				}`
	err := json.Unmarshal([]byte(jsonStr), &object)
	if err != nil {
		panic(err)
	}
	jsonParse("Z", object)
}

// jsonParse 逐层展开json
func jsonParse(prefix string, object interface{}) {
	switch object := object.(type) {
	case string:
		fmt.Printf("%s: %s\n", prefix, object)
	case []interface{}:
		for i, item := range object {
			jsonParse(fmt.Sprintf("%s.%d", prefix, i), item)
		}
	case map[string]interface{}:
		for key, value := range object {
			jsonParse(fmt.Sprintf("%s.%s", prefix, key), value)
		}
	default:
		fmt.Printf("%s: unhandle type %T\n", prefix, object)
	}
}
