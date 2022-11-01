package main

import "fmt"

func main() {
	fmt.Println("Введите сумму снятия со счёта:")
	var check_out int
	fmt.Scan(&check_out)
	if check_out%100 == 0 && check_out <= 100000 {
		fmt.Println("Операция успешно выполнена")
		fmt.Println("Вы сняли ", check_out, "рублей")
	} else if check_out%100 != 0 {
		fmt.Println("В банкомате доступны только купюры номиналом по 100 рублей")
	} else if check_out > 100000 {
		fmt.Println("Превышен лимит снятия")
	}

}
