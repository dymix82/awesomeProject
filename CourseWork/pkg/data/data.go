package data

import (
	"encoding/json"
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

// Находим максимальный ID чтобы присвоить его счетчику при запуске
func MaxId(v map[uint]*City) uint {
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

// Импортируем базу данных из CSV
func ImportCSV() {
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
	}
	Cid = MaxId(Storage)
}

// Грузим базу данных из json
func LoadDB() {
	in, err := os.ReadFile("database.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(in, &Storage)
	Cid = MaxId(Storage)
}

// Записываем базу данных из json
func SaveDB() {
	database, err := json.Marshal(Storage)
	err = os.WriteFile("database.json", database, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
