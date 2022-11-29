// Что нужно сделать  Разработайте программу, позволяющую ввести 10 целых чисел,
// а затем вывести из них количество чётных и нечётных чисел.
// Для ввода и подсчёта используйте разные циклы.
// Чтобы невводить числа наполняю его случайными через функцию
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func chetnechet(a [10]int) (chet, nechet int) { // Функция подсчета количества четных и нечетных чисел
	var lenght = len(a)
	for k := 0; k < lenght; k++ {
		if a[k]%2 == 0 {
			chet++
		} else {
			nechet++
		}
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
	fmt.Println("-------------------\nЗаполненный массив на 10 элементов\n------------------")
	a = fill(a)
	fmt.Println(a)
	chet, nechet := chetnechet(a)
	fmt.Printf("И в нем находится \n------------------\nЧетных %d чисел и %d нечетных", chet, nechet)

}
