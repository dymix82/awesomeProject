// Задание 2. Анонимные функции
//Что нужно сделать
// Напишите анонимную функцию, которая на вход получает массив типа integer,
// сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 6

func main() {
	array1 := make([]int, size)
	for k := 0; k < len(array1); k++ { // Наполняем массив
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)
		x := rand.Intn(20)
		//		fmt.Println(x)
		array1[k] = x
	}
	fmt.Println("Вводный массив: ", array1)
	array1Sorted := func(a []int) []int { // Анонимная функция, которая на вход получает массив типа integer,
		for i := 0; i < size-1; i++ { // сортирует его пузырьком
			for j := 0; j < size-i-1; j++ {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
		//	sort.Sort(sort.Reverse(sort.IntSlice(a))) Вот так можно сделать в одну строчку но это не спортивно
		con := 0
		for i := 0; i < (size)/2; i++ { // И переворачивает
			con = a[size-1-i]
			a[size-1-i] = a[i]
			a[i] = con
		}
		return a
	}
	fmt.Println("Отсортированный массив:", array1Sorted(array1))
}
