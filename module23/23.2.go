// Задание 2. Поиск символов в нескольких строках
// Что нужно сделать
// Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
// а возвращает 2D-массив,
// где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении
// i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:
// func parseTest(sentences []string, chars []rune)
package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}
	for i := range sentences {
		sentences[i] = strings.ToUpper(sentences[i])
	}

	x := parseTest(sentences, chars)
	fmt.Println("Сам 2D массив значений:")
	for i := range x {
		fmt.Println(x[i])
	}

}

func parseTest(sentences [4]string, chars [5]rune) [][]int { // Функция  возвращает 2D-массив где на позиции [i][j]
	var result [][]int            // стоит индекс вхождения символа j из chars
	for _, v := range sentences { // в последнее слово в предложении
		indexes := []int{}
		last := v[strings.LastIndex(v, " ")+1:] // Вычленяем последнее слово в приложении
		//	fmt.Println(last)
		for i := 0; i < len(chars); i++ { //
			j := strings.IndexRune(last, chars[i])
			indexes = append(indexes, j)
			if j != -1 {
				fmt.Println("в последнем слове приложения", v, "индекс буквы", string(chars[i]), "-", j)
			}
		}
		result = append(result, indexes)
	}
	return result
}
