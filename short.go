package main

import (
	"fmt"
	"strings"
)

// see: https://gist.github.com/dgritsko/9554733

var alphabet string = "abcdefghijklmnopqrstuvwxyz0123456789"
var base int = len(alphabet)

// codiert einen integer zu einem kurzen String
func encode(i int) string {

	if i == 0 {
		return string(alphabet[0])
	}

	s := ""
	for i > 0 {
		s += string(alphabet[i%base])
		i = i / base
	}
	return strings.Join([]string{reverse(s)}, "")

}

//decodiert einen kurzen String zu einem Integer
func decode(s string) int {

	i := 0

	//foreach (var c in s) {
	for _, char := range s {
		fmt.Printf("%s zu %s\n", s, string(char))
		i = (i * base) + strings.Index(alphabet, string(char))
		fmt.Printf("i ist: %d \n", i)
	}

	return i
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	fmt.Println("vim-go")

	/*
		for i := 0; i < 1000; i++ {

			fmt.Println("-------")
			fmt.Printf("kodiere: %d zu %s\n", i, encode(i))

			if decode(encode(1)) != i {
				fmt.Printf("%d is not %s\n", i, encode(i))
			}
		}
	*/
	fmt.Println("-------")
	fmt.Println("-------")
	fmt.Printf("kodiere %d zu %s\n", 23, encode(23))
	fmt.Printf("dekodiere %s zu %d\n", "ax", decode("ax"))
}
