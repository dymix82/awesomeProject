package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (u *User) addFriend(id int) error {
	if _, ok := storage[u.Id]; !ok {
		return errors.New("no such user")
	}
	u.Friends = append(u.Friends, storage[id].Id)
	storage[id].Friends = append(storage[id].Friends, u.Id)
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
	w.Write([]byte("User was created " + u.Name + "\n"))

	//	fmt.Println(storage)
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
	if _, ok := storage[src]; !ok {
		return errors.New("no such user")
	}
	if _, ok := storage[tgt]; !ok {
		return errors.New("no such user")
	}
	idx := slices.Index(storage[src].Friends, tgt)
	if idx == -1 {
		storage[tgt].addFriend(storage[src].Id)
		render.Status(r, http.StatusOK)
		fmt.Fprintf(w, "%s and %s are friends now \n", u.Name, storage[tgt].Name)
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
	if _, ok := storage[id]; !ok {
		return errors.New("no such user")
	}
	fmt.Fprintf(w, "%s is deleted\n", storage[id].Name)
	for _, vol := range storage[id].Friends {
		index := indexOf(id, storage[vol].Friends)
		removeIndex := RemoveIndex(storage[vol].Friends, index)
		storage[vol].Friends = removeIndex
	}
	delete(storage, id)
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
	if _, ok := storage[id]; !ok {
		return errors.New("no such user")
	}
	newAge, _ := strconv.Atoi(rep.Source)
	if newAge > 0 {
		storage[id].Age = newAge
		fmt.Fprintf(w, "Age of user %v is update to %+v\n", id, newAge)
	} else {
		return errors.New("something's wrong with his age")
	}

	render.Status(r, http.StatusOK)
	return nil
}
