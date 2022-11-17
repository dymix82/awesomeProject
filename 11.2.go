package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a10 10 20b 20 30c30 30 dd"
	words := strings.Split(str, " ")

	for _, i := range str {
		fmt.Println(words)
		i++
	}
	fmt.Println(words)
}
