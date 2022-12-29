package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
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

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

var u *User
var storage map[int]*User

func main() {
	storage = make(map[int]*User)
	r := chi.NewRouter()
	r.Method("GET", "/GetAll", Handler(GetAll))
	r.Method("POST", "/create", Handler(post))
	r.Method("GET", "/friends/", Handler(listFriends))
	r.Method("POST", "/make_friends", Handler(makeFriends))
	r.Method("DELETE", "/user", Handler(deleteUser))
	http.ListenAndServe(":8080", r)
}
func (u *User) addFriend(id int) error {
	if _, ok := storage[id]; !ok {
		return errors.New("no such user")
	}
	u.Friends = append(u.Friends, storage[id].Id)
	storage[id].Friends = append(storage[id].Friends, u.Id)
	return nil
}

func GetAll(w http.ResponseWriter, r *http.Request) error {
	bs, _ := json.Marshal(storage)
	//	fmt.Println(string(bs))
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
	//render.DefaultResponder(w, r, storage)
	return nil
}

func listFriends(w http.ResponseWriter, r *http.Request) error {
	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		return errors.New(idQuery)
	}
	id, err := strconv.Atoi(idQuery)
	lS := len(storage)
	log.Debugln(lS, id)
	if err != nil || id > lS {
		return errors.New(idQuery)
	}
	Fr, _ := json.Marshal(storage[id].Friends)
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
	fmt.Fprintf(w, "Person: %+v\n", p)
	data := &p
	src, _ := strconv.Atoi(data.Source)
	u, _ := storage[src]
	tgt, _ := strconv.Atoi(data.Target)
	storage[tgt].addFriend(storage[src].Id)
	render.Status(r, http.StatusOK)
	fmt.Fprintf(w, "%s i %s are friends now \n", u.Name, storage[tgt].Name)
	return nil
}
func deleteUser(w http.ResponseWriter, r *http.Request) error {
	var p delRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	data := &p
	id, _ := strconv.Atoi(data.Source)
	// fmt.Fprintf(w, "Person: %+v\n and number is %d", p, src)
	// delete(storage, id)
	for f := range u.Friends {
		removeIndex := RemoveIndex(storage[f].Friends, id)
		fmt.Fprintf(w, "Person: %+v\n  ", removeIndex)
	}
	render.Status(r, http.StatusOK)
	fmt.Fprintf(w, "%s is deleted\n", storage[id].Name)
	return nil
}
