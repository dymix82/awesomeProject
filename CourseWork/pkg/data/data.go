package data

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type City struct { // Описываем структуру
	Id         uint   `csv:"id",json:"id"`
	Name       string `csv:"name",json:"name"`
	Region     string `csv:"region",json:"region"`
	District   string `csv:"disct",json:"disct"`
	Population uint   `csv:"population",json:"population"`
	Foundation uint16 `csv:"found",json:"found"`
}

var Storage map[uint]*City

func ImportCSV() {
	Storage = make(map[uint]*City)
	Cities := []*City{}
	in, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	if err := gocsv.UnmarshalFile(in, &Cities); err != nil {
		panic(err)
	}
	for i, _ := range Cities {
		id := Cities[i].Id
		Storage[id] = Cities[i]
		fmt.Println(Storage[id].Name)
	}

}
