// Задание 1. Чётные и нечётные
// Что нужно сделать
//
//	Напишите функцию, которая принимает массив чисел, а возвращает два массива:
//	один из чётных чисел, второй из нечётных.
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const size = 10

func chetnechet(a []int) (chet []int, nechet []int) { // Функция разделения на два массива из четных и нечетных чисел
	var lenght = len(a)
	for k := 0; k < lenght; k++ {
		if a[k]%2 == 0 {
			chet = append(chet, a[k])
		} else {
			nechet = append(nechet, a[k])
		}
	}
	sort.Ints(chet)   // Сортируем для красоты
	sort.Ints(nechet) //
	return
}

func main() {
	var input_arr = make([]int, size)
	for k := 0; k < size; k++ { // Наполняем массив
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2) // Ждем чтобы массив не наполнился одинаковыми значениями
		x := rand.Intn(30)
		input_arr[k] = x
	}
	fmt.Printf("Исходный массив:\n%v\n", input_arr)
	fmt.Println("Массивы из его четных и не четных элементов элементов:")
	fmt.Println(chetnechet(input_arr))
}
