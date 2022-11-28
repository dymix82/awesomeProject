package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var a int

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	for i := 0; i <= 10; i++ {
		a += i * i
		log.Infof("i=%d a=%d", i, a)
	}
}
