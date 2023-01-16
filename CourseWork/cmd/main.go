package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"main/pkg/config"
	"main/pkg/data"
	"main/pkg/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

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
	r.Method("GET", "/get/", Handler(handlers.GetCitybyID)) // Вывод всех пользователей для дебага
	r.Method("POST", "/create", Handler(handlers.NewCity))  // Создание пользоватей
	//	r.Method("GET", "/cities/", Handler(user.ListCities))           // Вывод всех друзей
	//	r.Method("GET", "/make_friends", Handler(user.ListCitiesbyReg)) // Обработчик запросов в дружбу
	r.Method("DELETE", "/city", Handler(handlers.DeleteCity)) // Удаление пользователя
	r.Method("PUT", "/{id}", Handler(handlers.UpdatePop))     // Обновление возроста
	http.ListenAndServe(":"+config.Con.Apport, r)
	gr := make(chan struct{}, 1)
	graceful(gr)
	<-gr
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
