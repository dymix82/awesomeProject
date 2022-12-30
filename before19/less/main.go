package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
)

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
type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(500)
		w.Write([]byte("empty or invalid id"))
	}
}

func main() {
	r := chi.NewRouter()

	http.ListenAndServe(":8081", r)
	mux := http.NewServeMux()
	srv := service{make(map[int]*User)}
	r.Method("GET", "/get")
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.Get)
	mux.HandleFunc("/delete", srv.Delete)
	http.ListenAndServe("0.0.0.0:8080", mux)

}
func (u *User) toString() string {
	return fmt.Sprintf("User id is %d and his name is %s  and age is %d and %v  \n", u.Id, u.Name, u.Age, u.Friends)
}
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		uid += 1
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		s.store[uid] = &u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + u.Name + "\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
func (s *service) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.store {
			response += user.toString()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}

	}
	w.WriteHeader(http.StatusBadRequest)
}
func (s *service) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("User was deleted " + u.Name + "\n"))
		delete(s.store, u.Id)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
func (s *service) CustomHandler(w http.ResponseWriter, r *http.Request) error {
	response := ""
	for _, user := range s.store {
		response += user.toString()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}
