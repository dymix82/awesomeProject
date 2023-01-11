package main

import (
	"backend/pkg/db"
	"backend/pkg/user"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("empty or invalid id\n"))
	}
}

func main() {
	//	storage = make(map[int]*User)

	r := chi.NewRouter()
	//	r.Method("GET", "/GetAll", Handler(GetAll))             // Вывод всех пользователей для дебага
	r.Method("POST", "/create", Handler(post))                   // Создание пользоватей
	r.Method("GET", "/friends/", Handler(user.ListFriends))      // Вывод всех друзей
	r.Method("POST", "/make_friends", Handler(user.MakeFriends)) // Обработчик запросов в дружбу
	r.Method("DELETE", "/user", Handler(user.DeleteUser))        // Удаление пользователя
	r.Method("PUT", "/{id}", Handler(user.UpdateAge))            // Обновление возроста
	http.ListenAndServe(":8080", r)
}

func post(w http.ResponseWriter, r *http.Request) error {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	defer r.Body.Close()
	var u db.User
	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	db.Connect()
	_, err = db.DB.Exec("INSERT INTO Users (Users) VALUES($1)", &u)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	w.Write([]byte("User was created " + u.Name + "\n"))
	render.Status(r, http.StatusCreated)
	db.Close()
	return nil
}
