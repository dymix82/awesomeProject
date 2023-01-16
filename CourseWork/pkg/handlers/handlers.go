package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/ioutil"
	"main/pkg/data"
	"net/http"
	"strconv"
)

type AgeChange struct {
	Source string `json:"new pop"`
}
type delRequest struct {
	Source string `json:"source_id"`
}

func NewCity(w http.ResponseWriter, r *http.Request) error {
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
func GetCitybyID(w http.ResponseWriter, r *http.Request) error {
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
func UpdatePop(w http.ResponseWriter, r *http.Request) error {
	var p AgeChange
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	rep := &p
	idString := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idString)
	if _, ok := data.Storage[uint(id)]; !ok {
		return errors.New("no such city")
	}
	newPop, _ := strconv.Atoi(rep.Source)
	if newPop > 0 {
		data.Storage[uint(id)].Population = uint(newPop)
	} else {
		return errors.New("something's wrong with his age")
	}

	render.Status(r, http.StatusOK)
	return nil
}

func DeleteCity(w http.ResponseWriter, r *http.Request) error {
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
