package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Введите высоту елочки")
	var hight, count int
	fmt.Scan(&hight)
	star := "*"
	space := " "
	n := hight
	for i := 1; i <= hight; i++ {
		n--
		if i == 1 {
			line1 := strings.Repeat(space, n) + strings.Repeat(star, i)
			fmt.Println(line1)
			count++
		}
		if i > 1 {
			line1 := strings.Repeat(space, n) + strings.Repeat(star, i+count)
			fmt.Println(line1)
			count++
		}
	}
}
