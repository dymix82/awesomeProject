package main

import "fmt"

func main() {
	fmt.Println("введите количество покупателей")
	var n int
	var five, ten, twenty int
	var status bool
	fmt.Scan(&n)
	bills := make([]int, n)
	fmt.Println("введите их купюры 5,10 или 20")
	for i := 0; i < n; i++ {
		fmt.Scan(&bills[i])

	}
	//	fmt.Println(bills)
	for i := range bills {
		a := bills[i]
		//		fmt.Println(a)
		switch a {
		case 5:
			five++
			status = true
		case 10:
			ten++
			if five >= 1 {
				five--
				status = true
			} else {
				status = false
				goto label
			}
		case 20:
			twenty++
			if ten >= 1 && five >= 1 {
				ten--
				five--
				status = true
			} else {
				status = false
				goto label
			}
		default:
			fmt.Println("Введите только купюры 5,10,20")
			status = false
			goto label
		}
	}
label:
	fmt.Println("расчет сдачи")
	fmt.Println("Ввод:", bills)
	fmt.Println(status)
}
