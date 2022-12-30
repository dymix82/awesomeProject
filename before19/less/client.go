package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(500)
		w.Write([]byte("empty or invalid id"))
	}
}

var uid int

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}
type service struct {
	store map[int]*User
}

var s *service
var u *User
var storage map[int]*User

func main() {
	storage = make(map[int]*User)
	r := chi.NewRouter()
	r.Method("GET", "/GetAll", Handler(GetAll))
	r.Method("POST", "/create", Handler(post))
	http.ListenAndServe(":8080", r)
}
func (u *User) toString() string {
	return fmt.Sprintf("User id is %d and his name is %s  and age is %d and %v  \n", u.Id, u.Name, u.Age, u.Friends)
}
func GetAll(w http.ResponseWriter, r *http.Request) error {

	bs, _ := json.Marshal(storage)
	fmt.Println(string(bs))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bs))
	return nil
}
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
	w.Write([]byte("User was created " + u.Name + "\n"))

	fmt.Println(storage)
	render.Status(r, http.StatusCreated)
	//	render.DefaultResponder(w, r, storage)
	return nil
}
