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
main

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
	for {
		fmt.Println("Введите сообщение")
		count++
		var str string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str = scanner.Text()

		var filename string
		filename = "file.txt"
		num := strconv.Itoa(count)
		date := time.Now()
		datestr := date.Format("2006-01-02 03:04:05")
		str1 := num + " " + datestr + " " + str
		ioutil.WriteFile(filename, []byte(str1), 0666)


		if str == "exit" {
			f, err := ioutil.ReadFile("file.txt")
			if err != nil {
				log.Fatal(err)

			}
			fmt.Printf("Содержимое файла: %s\n", f)
		}
	}
}
