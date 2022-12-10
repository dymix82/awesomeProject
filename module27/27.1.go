package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct { // Описываем структуру
	name string

	age int

	grade int
}

func newStudent(name string, age int, grade int) *Student { //Функция добавления данных в структуру
	p := Student{name: name,
		age:   age,
		grade: grade}
	return &p
}
func main() {
	StudentName := make(map[int]*Student)
	fmt.Println("Введите студентов построчно в формате имя возраст курс через пробел, курс от 1 до 5")
	index := 1
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		s := strings.Split(text, " ")
		if err == io.EOF { // Выходим и печатаем список студентов
			fmt.Println("Студенты из хранилища:")
			for _, value := range StudentName {
				fmt.Println(value.name, value.age, value.grade)
			}
			break
		}
		if len(s) != 3 {
			fmt.Println("Ошибка при вводе, попробуйте еще раз")
			continue
		}
		age, err2 := strconv.Atoi(s[1])
		if err2 != nil || age <= 0 {
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
		StudentName[index] = newStudent(s[0], age, grade) // Добавляем функцией студента
		fmt.Println("Студент добавлен")
		index++
	}
}
