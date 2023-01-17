package data

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
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
var Cid uint

func init() {
	Storage = make(map[uint]*City)
}
func MaxId(v map[uint]*City) uint { // Находим максимальный ID чтобы присвоить его счетчику при запуске
	keys := make([]uint, 0, len(Storage))
	for key := range Storage {
		keys = append(keys, key)
	}
	for i := 0; i <= len(keys)-1; i++ {
		for j := 0; j < len(keys)-1-i; j++ {
			if keys[j] > keys[j+1] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}
	return keys[len(v)-1]
}
func ImportCSV() { // Импортируем базу данных из CSV
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
	Cid = MaxId(Storage)
}

func LoadDB() { // Грузим базу данных из json
	in, err := os.ReadFile("database.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(in, &Storage)
	Cid = MaxId(Storage)
}

func SaveDB() { // Записываем базу данных из json
	database, err := json.Marshal(Storage)
	err = os.WriteFile("database.json", database, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
