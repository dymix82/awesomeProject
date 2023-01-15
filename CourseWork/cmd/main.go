package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/go-chi/chi/v5"
	"main/pkg/config"
	"main/pkg/data"
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

func main() {
	importCSV := flag.Bool("import", false, "a bool")
	flag.Parse()
	if *importCSV {
		data.ImportCSV()
	}
	r := chi.NewRouter()
	r.Method("GET", "/get/", Handler(getCitybyID)) // Вывод всех пользователей для дебага
	//	r.Method("POST", "/create", Handler(post))                      // Создание пользоватей
	//	r.Method("GET", "/cities/", Handler(user.ListCities))           // Вывод всех друзей
	//	r.Method("GET", "/make_friends", Handler(user.ListCitiesbyReg)) // Обработчик запросов в дружбу
	// r.Method("DELETE", "/user", Handler(user.CityDelete))           // Удаление пользователя
	// r.Method("PUT", "/{id}", Handler(user.UpdatePop))               // Обновление возроста
	http.ListenAndServe(":"+config.Con.Apport, r)
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
		return nil
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("empty or invalid id\n"))
		return nil
	}
	return nil
}

//func graceful(ch chan<- struct{}, srv *http.Server, s *data.City) {
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
//	go func() {
//		<-c
//		log.Println("Shutting down...")
//		if err := srv.Shutdown(context.Background()); err != nil {
//			log.Printf("Server shutdown error: %s\n", err)
//		}
//		//	saveDb(s)
//		close(ch)
//	}()
//}
