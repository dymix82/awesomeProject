package main

import (
	"bufio"
	"fmt"
	"io"
	"module28/pkg/storage"
	"module28/pkg/student"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите студентов построчно в формате имя возраст курс через пробел, курс от 1 до 5")
	st := storage.New()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		s := strings.Split(text, " ")
		if err == io.EOF { // Выходим и печатаем список студентов
			for n := range st {
				s := st.Get(n)
				fmt.Println(s.Name(), s.Age(), s.Grade())
			}
			break
		}
		if len(s) != 3 {
			fmt.Println("Ошибка при вводе, попробуйте еще раз")
			continue
		}
		Age, err2 := strconv.Atoi(s[1])
		if err2 != nil || Age <= 0 {
			fmt.Println("Ошибка при вводе возроста, введите строку еще раз")
			fmt.Println("Введите студентов построчно в формате имя возраст курс через пробел, курс от 1 до 5")
			continue
		}
		grade, err3 := strconv.Atoi(s[2])
		if err3 != nil {
			fmt.Println("Ошибка при вводе курса, введите строку еще раз")
			fmt.Println("Введите студентов построчно в формате имя возраст курс через пробел, курс от 1 до 5")
			continue
		}
		if grade < 0 || grade > 5 {
			fmt.Println("Ошибка при вводе курса, введите строку еще раз")
			fmt.Println("Введите студентов построчно в формате имя возраст курс через пробел, курс от 1 до 5")
			continue
		}
		sm := student.New(s[0], Age, grade)
		st.Put(sm)
		fmt.Println("Студент добавлен")
		//	index++
	}
}
