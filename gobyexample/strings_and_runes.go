package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "哈哈"

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("%#U starts at %d\n", r, i)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", r, i)
		w = width

		examineRune(r)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '哈' {
		fmt.Println("found 哈")
	}
}
