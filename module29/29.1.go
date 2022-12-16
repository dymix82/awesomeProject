package main

import (
	"fmt"
	"strconv"
)

var input string
var sqr int

func main() {
	ch1 := make(chan int)
	fmt.Println("Введите число , если хотите остановиться введите 'стоп'")
	for {
		fmt.Scan(&input)
		if input == "стоп" {
			break
		}
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Это не число, попробуйте ввести число еще раз")
			continue
		}
		go sqrt(value, ch1)
		value2 := <-ch1
		fmt.Println(value2)
		go multiplyByTwo(value2, ch1)
		fmt.Println(<-ch1)
	}

}

func multiplyByTwo(num int, out chan<- int) {
	result := num * 2
	out <- result
}
func sqrt(num int, out chan<- int) {
	result := num * num
	out <- result
}
