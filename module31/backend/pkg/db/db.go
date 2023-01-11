package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	host     = "host.docker.internal"
	port     = 49153
	user     = "db_user"
	password = "jw8s0F4"
	dbname   = "friendsdb"
)

var DB *sqlx.DB

type User struct {
	//	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

func (a User) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *User) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	con, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	DB = con
}

func Close() {
	DB.Close()
}
