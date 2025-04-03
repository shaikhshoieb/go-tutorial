package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "hello, 世界"
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, w := 0, 0; i < len(s); i += w {
		// fmt.Printf("%x ", s[i])
		runeVal, width := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%#U \n", runeVal)
		fmt.Printf("width: %d ", width)
		w = width
	}
}
