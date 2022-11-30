// Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
// Заполните упорядоченный массив из 12 элементов и введите число.
// Необходимо реализовать поиск первого вхождения заданного числа в массив.
// Сложность алгоритма должна быть минимальная.
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const size2 = 12

func searchIndex(arr []int, b int) int { // Функция поиска
	result := -1
	min := 0
	max := size2 - 1
	for max >= min {
		middle := (max + min) / 2 // находим середину значений
		if arr[middle] == b {     // Если значение лежит в середине
			result = middle
			break // отдаем результат и прерываемся
		} else if arr[middle] > b { // Если значение больше сдвигаем максимум
			max = middle - 1
		} else { // Сдвигаем минимум
			min = middle + 1
		}
	}
	return result
}
func main() {
	var a int
	var ArraySorted = make([]int, size2)
	for k := 0; k < size2; k++ { // Наполняем массив
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2) // Ждем чтобы массив не наполнился одинаковыми значениями
		x := rand.Intn(size2)
		ArraySorted[k] = x
	}
	sort.Ints(ArraySorted) // Сортируем
	fmt.Println(ArraySorted)
	fmt.Println("Введите число для поиска:")
	fmt.Scan(&a)
	Result := searchIndex(ArraySorted, a)
	if Result == -1 { // Если числа нет
		fmt.Printf("Числа %d в массиве нет , а следовательно выводим: %d", a, Result)
	} else { // Иначе выводим результат
		fmt.Printf("Индекс первого вхождения заданного числа %d в массив - %d", a, Result)
	}
}
