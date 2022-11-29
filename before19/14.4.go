// Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
//
//	Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
//	Затем вызовите по очереди три функции, передавая результат из одной в другую.
package main

import "fmt"

var (
	a = 10
	b = 20
	c = 30
)

func main() {
	fmt.Println(firstfunc(secondfunc(thirdfunc(10))))
}

func firstfunc(x int) int {
	return a + b + x
}
func secondfunc(y int) int {
	return c + b + y
}
func thirdfunc(z int) int {
	return c + a + z
}
