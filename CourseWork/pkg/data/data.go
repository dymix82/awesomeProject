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
	Cid = 1100
	Storage = make(map[uint]*City)
}
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
		fmt.Println(Storage[id].Name)
	}

}

func LoadDB() {
	in, err := os.ReadFile("database.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(in, &Storage)
}

func SaveDB() {
	database, err := json.Marshal(Storage)
	err = os.WriteFile("database.json", database, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//jsonFile, err := os.ReadFile("database.json")
	//
	//	if err != nil {
	//		panic(err)
	//	}

	//	jsonData, err := json.Marshal(Storage)
	//	jsonFile.Write(jsonData)
	//	jsonFile.Close()
	//	fmt.Println("JSON data written to ", jsonFile.Name())
}

//func FindbyReg(string) {
//	for i, _ := range Storage {
//		if v, found := m["pi"]; found {
//			fmt.Println(v)
//		}
//	}
//}
