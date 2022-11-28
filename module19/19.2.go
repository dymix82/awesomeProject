//Задание 2. Сортировка пузырьком
//Что нужно сделать
//Отсортируйте массив длиной шесть пузырьком.
//
//Советы и рекомендации
//Принцип сортировки пузырьком можно посмотреть на «Википедии», там есть наглядная демонстрация, или на YouTube.
//В качестве среды разработки может помочь GoLand или VS Code.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 6

var array1 [size]int

func fill(a [size]int) (b [size]int) { // Функция наполнения массива
	for k := 0; k < len(a); k++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)
		x := rand.Intn(20)
		//		fmt.Println(x)
		b[k] = x
	}
	return
}
func BubbleSort(a [size]int) [size]int {
	for i := 0; i < size-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	array1 = fill(array1)
	array1_sorted := BubbleSort(array1)
	fmt.Println("Вводный массив: ", array1)
	fmt.Println("Отсортированный массив:", array1_sorted)
}
