package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите IQ космонавта  1")
	var a int
	fmt.Scan(&a)
	fmt.Println("Введите IQ космонавта  2:")
	var b int
	fmt.Scan(&b)
	fmt.Println("Введите IQ космонавта  3:")
	var c int
	fmt.Scan(&c)
	var max int
	if a < c && b < c {
		max = c
		fmt.Print("Капитан коробля стал космонафт №3 с IQ: ", max)
	} else if b < a && c < a {
		max = a
		fmt.Print("Капитан коробля стал космонафт №1 с IQ: ", max)
	} else if c < b && a < b {
		max = b
		fmt.Print("Капитан коробля стал космонафт №2 с IQ: ", max)
	}
	// fmt.Println(min, mid, max)

}
