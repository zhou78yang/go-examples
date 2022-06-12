package main

import (
	"fmt"
)

func main() {
	s := "Go语言"

	fmt.Println(len(s), len([]byte(s)), len([]rune(s))) // 8 8 4

	// 0 G 1 o 2 语 5 言
	for i, r := range s {
		fmt.Printf("%d %c ", i, r)
	}
	fmt.Println()
}
