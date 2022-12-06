// Цель задания
// Написать программу для нахождения подстроки в кириллической подстроке.
// Программа должна запускаться с помощью команды:
// go run main.go --str "строка для поиска" --substr "поиска"
// Для реализации такой работы с флагами воспользуйтесь пакетом flags,
// а для поиска подстроки в строке вам понадобятся руны.
package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string
	var substr string
	flag.StringVar(&str, "str", "null", "str")
	flag.StringVar(&substr, "substr", "null", "substr")
	flag.Parse()
	strRunes := []rune(str)
	substrRunes := []rune(substr)
	fmt.Println(stringSearch(strRunes, substrRunes))
}
func stringSearch(strRunes []rune, substrRunes []rune) bool {
	j := 0
	for _, r := range strRunes {
		if r == substrRunes[j] {
			j++
			if j == len(substrRunes) {
				return true
			}
		} else {
			j = 0
		}
	}
	return false
}
