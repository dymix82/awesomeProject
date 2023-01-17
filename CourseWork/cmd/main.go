package main

import (
	"context"
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
	"time"
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
		if _, err := os.Stat("database.json"); err == nil {
			data.LoadDB()
		} else {
			data.ImportCSV()
		}
	}
	r := chi.NewRouter()
	r.Method("GET", "/get/", Handler(handlers.GetCityby))       // Вывод городов в соотвествии с запросом
	r.Method("POST", "/create", Handler(handlers.NewCity))      // Добовление городов
	r.Method("DELETE", "/delete", Handler(handlers.DeleteCity)) // Удаление городов
	r.Method("PUT", "/{id}", Handler(handlers.UpdatePop))       // Изменение населения города
	srv := &http.Server{
		Addr:    ":" + config.Con.Apport,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		data.SaveDB()
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
