// Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
//
//	Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
//	Затем вызовите по очереди три функции, передавая результат из одной в другую.
package main

import "fmt"

var input int

func main() {
	fmt.Println("----------\nВведите число\n----------")
	fmt.Scan(&input)
	fmt.Println(thirdfunc(input))
}

func firstfunc(a int) int {
	return a + input
}
func secondfunc(b int) int {
	return firstfunc(b) + input
}
func thirdfunc(c int) int {
	return secondfunc(c) + input
}
