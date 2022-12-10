package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string

	age int

	grade int
}

func newStudent(name string, age int, grade int) Student {
	p := Student{name: name,
		age:   age,
		grade: grade}
	return p
}
func main() {
	StudentName := make(map[int]Student)
	fmt.Println("Введите студенов построчно в формате имя возраст оценку через пробел, оценку от 1 до 5")
	index := 1
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		//	fmt.Println(text)
		s := strings.Split(text, " ")
		//	for _, v := range s {
		//		fmt.Println(v)
		//	}
		if err == io.EOF {
			fmt.Println(StudentName)
			goto label
		}
		age, err2 := strconv.Atoi(s[1])
		if err2 != nil {
			fmt.Println("Ошибка при вводе возроста, введите строку еще раз")
			fmt.Println("Введите студенов построчно в формате имя возраст оценку через пробел, оценку от 1 до 5")
			continue
		}
		grade, err3 := strconv.Atoi(s[2])
		if err3 != nil {
			fmt.Println("Ошибка при вводе оценки, введите строку еще раз")
			fmt.Println("Введите студенов построчно в формате имя возраст оценку через пробел, оценку от 1 до 5")
			continue
		}
		if grade < 0 || grade > 5 {
			fmt.Println("Ошибка при вводе оценки, введите строку еще раз")
			fmt.Println("Введите студенов построчно в формате имя возраст оценку через пробел, оценку от 1 до 5")
			continue
		}
		//	print(s[2])
		//	sp := Student{name: s[0], age: age, grade: grade}
		StudentName[index] = newStudent(s[0], age, grade)
		fmt.Println("Студент добавлен")
		index++
	}
label:
}
