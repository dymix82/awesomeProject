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

//	func newPerson(name string) *person {
//		p := person{name: name}
//		p.age = 42
//		return &p
//	}
func main() {
	StudentName := make(map[int]Student)
	fmt.Println("Введите студенов построчно в формате имя возраст курс  через пробел")
	//	stdin, err := io.ReadAll(os.Stdin)
	//	if err != nil {
	//
	//	}
	//	str := string(stdin)
	//	words := strings.Fields(str)
	//	for i,v:= range words {
	//		if
	//	}
	//
	//	fmt.Println(words)
	index := 1
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		//	fmt.Println(text)
		s := strings.Split(text, " ")
		if err == io.EOF {
			fmt.Println(StudentName)
			goto label
		}
		age, _ := strconv.Atoi(s[1])
		grade, _ := strconv.Atoi(s[2])
		print(s[2])
		sp := Student{name: s[0], age: age, grade: grade}
		StudentName[index] = sp
		fmt.Println(StudentName)
		index++
	}
label:
}
