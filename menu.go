package main

import "fmt"

func main() {
	fmt.Println("Введите сумму заказа:")
	var check_sum int
	fmt.Scan(&check_sum)
	fmt.Println("Введите день недели (от 1 до 7):")
	var dayOfweek int
	fmt.Scan(&dayOfweek)
	if dayOfweek >= 1 && dayOfweek <= 5 {
		if check_sum >= 10000 {
			check_sum /= 10
			fmt.Println("Сумма скидки составила:", check_sum)
		} else {
			fmt.Println("Суммы чека меньше необходимой для ссылки ")
		}
	} else if dayOfweek <= 7 && dayOfweek >= 6 {
		fmt.Println("Скидка в выходные дни не предоставляется")

	} else {
		fmt.Println("Введите корректный день недели")
	}

}
