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

	s := ""
	for i > 0 {
		s += string(alphabet[i%base])
		i = i / base
		//fmt.Printf("s  = %#v \n", s)
	}
	if i == 0 {
		s += string(alphabet[0])
	}

	result := strings.Join([]string{reverse(s)}, "")
	//fmt.Printf("result  = %#v \n", result)
	return result
}

//decodiert einen kurzen String zu einem Integer
func decode(s string) int {

	i := 0

	//foreach (var c in s) {
	for _, char := range s {
		//fmt.Printf("char  = %#v \n", char)

		//fmt.Printf("%s zu %s\n", s, string(char))
		i = (i * base) + strings.Index(alphabet, string(char))

		//fmt.Printf("i  = %#v \n", i)
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

	//fmt.Printf("base  = %#v \n", base)
	for i := 0; i < 9000000; i++ {

		fmt.Println("-------")
		fmt.Printf("kodiere: %d zu %s\n", i, encode(i))

		if decode(encode(i)) != i {
			fmt.Printf("%d is not %s\n", i, encode(i))
		}
	}
	/*
		fmt.Println("---ENCODE ---")
		fmt.Printf("kodiere %d zu %#v\n", 23, encode(23))
		fmt.Println("--- DECODE ---")
		fmt.Printf("dekodiere %s zu %d\n", "ax", decode("ax"))
		fmt.Println("-------")
	*/
}
