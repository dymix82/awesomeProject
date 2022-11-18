// Задание 2. Функция, реверсирующая массив
// Что нужно сделать
// Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке по сравнению с исходным.
// Напишите программу, демонстрирующую работу этого метода.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func viceversa(a [10]int) (b [10]int) { // Функция реверсирующая массив
	var lenght = len(a)
	for k := 0; k < lenght; k++ {
		b[k] = a[(lenght-1)-k]
	}
	return
}
func fill(a [10]int) (b [10]int) { // Функция наполнения массива
	var lenght = len(a)
	for k := 0; k < lenght; k++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)
		x := rand.Intn(20)
		//		fmt.Println(x)
		b[k] = x
	}
	return
}

func main() {
	var a [10]int
	var b [10]int
	fmt.Println("-------------------\nЗаполненный массив на 10 элементов\n------------------")
	a = fill(a)
	fmt.Println(a)
	fmt.Println("-------------------\nТот же массив на 10 элементов, но задом наперед\n------------------")
	b = viceversa(a)
	fmt.Println(b)
}
