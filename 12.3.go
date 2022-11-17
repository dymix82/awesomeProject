package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	count := 0
	var filename string
	filename = "file.txt"
	ioutil.WriteFile(filename, []byte(""), 0666)
	for {
		fmt.Println("Введите сообщение")
		count++
		var str string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str = scanner.Text()

		num := strconv.Itoa(count)
		date := time.Now()
		datestr := date.Format("2006-01-02 03:04:05")
		str1 := num + " " + datestr + " " + str + "\n"
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(str1)); err != nil {
			log.Fatal(err)
		}
		if str == "exit" {
			f, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)

			}
			fmt.Printf("Содержимое файла: %s\n", f)
		}
		f.Close()
	}
}
