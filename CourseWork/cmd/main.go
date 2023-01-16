package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"main/pkg/config"
	"main/pkg/data"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type delRequest struct {
	Source string `json:"source_id"`
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("empty or invalid id\n"))
	}
}

func main() {

	importCSV := flag.Bool("import", false, "a bool")
	flag.Parse()
	if *importCSV {
		data.ImportCSV()
	} else {
		data.LoadDB()
	}
	r := chi.NewRouter()
	r.Method("GET", "/get/", Handler(getCitybyID)) // Вывод всех пользователей для дебага
	r.Method("POST", "/create", Handler(newCity))  // Создание пользоватей
	//	r.Method("GET", "/cities/", Handler(user.ListCities))           // Вывод всех друзей
	//	r.Method("GET", "/make_friends", Handler(user.ListCitiesbyReg)) // Обработчик запросов в дружбу
	r.Method("DELETE", "/city", Handler(deleteCity)) // Удаление пользователя
	// r.Method("PUT", "/{id}", Handler(user.UpdatePop))               // Обновление возроста
	http.ListenAndServe(":"+config.Con.Apport, r)
	gr := make(chan struct{}, 1)
	graceful(gr)
	<-gr
}
func newCity(w http.ResponseWriter, r *http.Request) error {
	data.Cid++
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	defer r.Body.Close()
	var u data.City
	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	u.Id = data.Cid
	data.Storage[u.Id] = &u
	w.Write([]byte("City was added " + u.Name + "\n"))
	Fr, _ := json.Marshal(data.Storage)
	w.Write([]byte(Fr))
	//	fmt.Println(storage)
	render.Status(r, http.StatusCreated)
	return nil
}
func getCitybyID(w http.ResponseWriter, r *http.Request) error {
	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		return errors.New(idQuery)
	}
	id, _ := strconv.Atoi(idQuery)
	if _, ok := data.Storage[uint(id)]; ok {
		Fr, _ := json.Marshal(data.Storage[uint(id)])
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(Fr))
		// data.SaveDB()
		return nil
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("empty or invalid id\n"))
		return nil
	}
	return nil
}
func deleteCity(w http.ResponseWriter, r *http.Request) error {
	var p delRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	req := &p
	id, _ := strconv.Atoi(req.Source)
	if _, ok := data.Storage[uint(id)]; !ok {
		return errors.New("no such user")
	}
	fmt.Fprintf(w, "%s is deleted\n", data.Storage[uint(id)].Name)
	//	for _, vol := range storage[id].Friends {
	//		index := indexOf(id, storage[vol].Friends)
	//		removeIndex := RemoveIndex(storage[vol].Friends, index)
	//		storage[vol].Friends = removeIndex
	//	}
	delete(data.Storage, uint(id))
	render.Status(r, http.StatusOK)
	return nil
}

func graceful(ch chan<- struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down...")
		data.SaveDB()
		close(ch)
	}()
}
