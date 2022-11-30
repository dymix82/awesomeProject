// Задание 1.
// Что нужно сделать
// Заполните массив неупорядоченными числами на основе генератора случайных чисел.
// Введите число. Программа должна найти это число в массиве и вывести, сколько
// чисел находится в массиве после введённого.
// При отсутствии введённого числа в массиве — вывести 0.
// Для удобства проверки реализуйте вывод массива на экран.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

var result int

func fill([size]int) (b [size]int) { // Функция наполнения массива
	for k := 0; k < size; k++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)
		x := rand.Intn(5 * size)
		//		fmt.Println(x)
		b[k] = x
	}
	return
}
func find(a [size]int, b int) int { // Функция поиска по  массиву
	result = 0
	for k := 0; k < size; k++ { // Цикл поиска
		if b == a[k] {
			result = size - k - 1
			if k == size-1 { // Если число в конце присваеваем индекс за массивом, чтобы не пересекался с 0
				result = 11 // Т.к по условию при ненахождении числа нужно выводить 0
				return result
			}
			break
		}
	}
	return result
}

func main() {
	var Array [size]int
	var Num int
	Array = fill(Array)
	fmt.Println(Array)
	fmt.Println("Введите число для поиска:")
	fmt.Scan(&Num)
	Res := find(Array, Num)
	if Res == 11 { // Если число в конце
		fmt.Println("Искомое число находится в конце массива и чисел после него нет")
	} else if Res != 0 { // Если число в массиве
		fmt.Printf("Количество чисел в массиве после введенного числа %d составляет %d", Num, Res)
	} else { // Если число не в массиве
		fmt.Printf("Данное число не найдено, а следовательно выводим %d", Res)
	}
}
