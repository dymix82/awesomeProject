package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	host     = "localhost"
	port     = 49153
	user     = "db_user"
	password = "jw8s0F4"
	dbname   = "friendsdb"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("empty or invalid id\n"))
	}
}

var uid int

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}
type FriendRequest struct {
	Source string `json:"source_id"`
	Target string `json:"target_id"`
}
type delRequest struct {
	Source string `json:"source_id"`
}
type AgeChange struct {
	Source string `json:"new age"`
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
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

var storage map[int]*User
var DB *sql.DB

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	con, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	DB = con
}

func Close() {
	DB.Close()
}

func checkUser(id int) bool {
	Connect()
	var result int
	check := `select 1 from users where id = $1 limit 1;`
	e := DB.QueryRow(check, id).Scan(&result)
	if e != nil {
		fmt.Println(e.Error())
	}
	if result == 1 {
		return true
	} else {
		return false
	}
}
func main() {
	storage = make(map[int]*User)

	r := chi.NewRouter()
	//	r.Method("GET", "/GetAll", Handler(GetAll))             // Вывод всех пользователей для дебага
	r.Method("POST", "/create", Handler(post))              // Создание пользоватей
	r.Method("GET", "/friends/", Handler(listFriends))      // Вывод всех друзей
	r.Method("POST", "/make_friends", Handler(makeFriends)) // Обработчик запросов в дружбу
	r.Method("DELETE", "/user", Handler(deleteUser))        // Удаление пользователя
	r.Method("PUT", "/{id}", Handler(updateAge))            // Обновление возроста
	http.ListenAndServe(":8080", r)
}
func UserFriendsToSlice(id int) []int {
	Connect()
	var result string
	listFriends := `select users->'friends' from users where id = $1;`
	e := DB.QueryRow(listFriends, id).Scan(&result)
	if e != nil {
		fmt.Println(e.Error())
	}

	if result != "null" {
		trimmed := strings.Trim(result, "[]")
		strings := strings.Split(trimmed, ", ")
		ints := make([]int, len(strings))
		for i, s := range strings {
			ints[i], _ = strconv.Atoi(s)
		}
		return ints
	} else {
		null := make([]int, 0)
		return null
	}
}
func (u *User) addFriend(id int) error {
	if checkUser(id) != true {
		return errors.New("no such user")
	}
	u.Friends = append(u.Friends, id)
	//	Connect()
	//	addFriendQ := `UPDATE users SET users = users || jsonb_build_object('friends', $1::int[]) WHERE id = $2`
	//	_, e := DB.Exec(addFriendQ, pq.Array(u.Friends), u.Id)
	//	if e != nil {
	//		fmt.Println(e.Error())
	//	}
	//	Close()
	storage[id].Friends = append(storage[id].Friends, u.Id)
	return nil
}
func MakeFriends(id1, id2 int) error {
	if checkUser(id1) != true {
		return errors.New("no such user")
	}
	if checkUser(id2) != true {
		return errors.New("no such user")
	}
	Connect()
	result1 := UserFriendsToSlice(id1)
	result2 := UserFriendsToSlice(id2)
	res1 := append(result2, id1)
	res2 := append(result1, id2)
	addFriendQ := `UPDATE users SET users = users || jsonb_build_object('friends', $1::int[]) WHERE id = $2`
	_, e := DB.Exec(addFriendQ, pq.Array(res1), id2)
	if e != nil {
		fmt.Println(e.Error())
	}
	_, e2 := DB.Exec(addFriendQ, pq.Array(res2), id1)
	if e2 != nil {
		fmt.Println(e.Error())
	}
	Close()
	fmt.Println(result1, result2)
	return nil
}

//	func GetAll(w http.ResponseWriter, r *http.Request) error {
//		bs, _ := json.Marshal(storage)
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(bs))
//		return nil
//	}
func post(w http.ResponseWriter, r *http.Request) error {
	uid += 1
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	defer r.Body.Close()
	var u User
	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	u.Id = uid
	storage[uid] = &u
	Connect()
	_, err = DB.Exec("INSERT INTO Users (Users) VALUES($1)", &u)
	if err != nil {
		log.Fatal(err)
	}
	Close()
	w.Write([]byte("User was created " + u.Name + "\n"))
	render.Status(r, http.StatusCreated)
	return nil
}

func listFriends(w http.ResponseWriter, r *http.Request) error {
	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		return errors.New(idQuery)
	}
	id, err := strconv.Atoi(idQuery)
	lS := len(storage)
	if err != nil || id > lS {
		return errors.New(idQuery)
	}
	out := make(map[int]*User)
	for i, vol := range storage[id].Friends {
		out[i] = storage[vol]
	}
	Fr, _ := json.Marshal(out)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Fr))
	return nil
}
func makeFriends(w http.ResponseWriter, r *http.Request) error {
	var p FriendRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	data := &p
	src, _ := strconv.Atoi(data.Source)
	u, _ := storage[src]
	tgt, _ := strconv.Atoi(data.Target)
	if checkUser(src) != true {
		return errors.New("no such user")
	}
	if checkUser(tgt) != true {
		return errors.New("no such user")
	}

	idx := slices.Index(storage[src].Friends, tgt)
	if idx == -1 {
		storage[tgt].addFriend(storage[src].Id)
		render.Status(r, http.StatusOK)
		fmt.Fprintf(w, "%s and %s are friends now \n", u.Name, storage[tgt].Name)
		MakeFriends(src, tgt)
		return nil
	} else {
		return errors.New("no such user")
	}
}
func deleteUser(w http.ResponseWriter, r *http.Request) error {
	var p delRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	req := &p
	id, _ := strconv.Atoi(req.Source)
	if checkUser(id) != true {
		return errors.New("no such user")
	}
	fmt.Fprintf(w, "%s is deleted\n", storage[id].Name)
	for _, vol := range storage[id].Friends {
		index := indexOf(id, storage[vol].Friends)
		removeIndex := RemoveIndex(storage[vol].Friends, index)
		storage[vol].Friends = removeIndex
	}
	delete(storage, id)
	Connect()
	deleteStmt := `delete from Users where id=$1`
	_, e := DB.Exec(deleteStmt, id)
	if e != nil {
		errors.New("no such user")
	}
	render.Status(r, http.StatusOK)
	return nil
}
func updateAge(w http.ResponseWriter, r *http.Request) error {
	var p AgeChange
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	rep := &p
	idString := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idString)
	if checkUser(id) != true {
		return errors.New("no such user")
	}
	newAge, _ := strconv.Atoi(rep.Source)
	if newAge > 0 {
		Connect()
		updateAge2 := `UPDATE users SET users = users || jsonb_build_object('age', $1::int) WHERE id = $2`
		_, e := DB.Exec(updateAge2, newAge, id)
		if e != nil {
			fmt.Println(e.Error())
		}
		Close()
		fmt.Fprintf(w, "Age of user %v is update to %+v\n", id, newAge)
	} else {
		return errors.New("something's wrong with his age")
	}

	render.Status(r, http.StatusOK)
	return nil
}
