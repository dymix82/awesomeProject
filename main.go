package main

import "fmt"

func main() {
	fmt.Println("расчет сдачи")
	fmt.Println("введите количество покупателей")
	var n int
	fmt.Scan(&n)
	bills := make([]int, n)
	fmt.Println("введите их купюры 5,10 или 20")
	for i := 0; i < n; i++ {
		fmt.Scan(&bills[i])
	}
	fmt.Println(bills)
	a := lemonadeChange(bills)
	fmt.Println(a)
}
func lemonadeChange(bills []int) bool {
	var five, ten, twenty int
	var status bool = false
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
			status = false
			goto label
		}
	}

	switch status {
	case status == true:
		return true
	case status == false:
		return false
	}
label:
	return false
}
