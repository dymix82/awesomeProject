package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введие зарплату 1")
	var a int
	fmt.Scan(&a)
	fmt.Println("Введие зарплату 2:")
	var b int
	fmt.Scan(&b)
	fmt.Println("Введие зарплату 3:")
	var c int
	fmt.Scan(&c)
	var min int
	var mid int
	var max int
	if a < b && b < c {
		min = a
		mid = b
		max = c
	} else if b < a && a < c {
		min = b
		mid = a
		max = c
	} else if c < a && a < b {
		min = c
		mid = a
		max = b
	} else if b < c && c < a {
		min = b
		mid = c
		max = a
	} else if c < b && b < a {
		min = c
		mid = b
		max = a
	} else if a < c && c < b {
		min = a
		mid = c
		max = b
	}
	fmt.Println(min, mid, max)
	fmt.Println("Pазница между самой большой и самой маленькой зарплатой:", max-min)
	fmt.Println("Средняя зарплата:", (max+mid+min)/3)
}
