package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Apport     string `yaml:"apport"`
	Dbhost     string `yaml:"dbhost"`
	Dbport     int    `yaml:"dbport"`
	Dbuser     string `yaml:"dbuser"`
	Dbpassword string `yaml:"dbpassword"`
	Dbname     string `yaml:"dbname"`
}

var DB *sqlx.DB
var Cfg Config
var Con *Config

type User struct {
	//	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

func GetConf(file string, cnf interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, cnf)
	}
	return err
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
func init() {

	if err := GetConf("conf.yml", &Cfg); err != nil {
		log.Panicln(err)
	}
	Con = &Cfg
}

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Con.Dbhost, Con.Dbport, Con.Dbuser, Con.Dbpassword, Con.Dbname)
	con, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	DB = con
}

func Close() {
	DB.Close()
}
