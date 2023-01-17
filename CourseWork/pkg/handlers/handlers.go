package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/ioutil"
	"main/pkg/data"
	"net/http"
	"strconv"
)

type PopChange struct {
	Source string `json:"new pop"`
}
type delRequest struct {
	Source string `json:"source_id"`
}

func NewCity(w http.ResponseWriter, r *http.Request) error { // добавление новой записи в список городов;
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
	render.Status(r, http.StatusCreated)
	return nil
}

func UpdatePop(w http.ResponseWriter, r *http.Request) error { // обновление информации о численности населения города по указанному id;
	var p PopChange
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
		return errors.New("Количиство жителей не может быть отрицательным")
	}
	w.Write([]byte("Population is changed in city: " + data.Storage[uint(id)].Name + "\n"))
	render.Status(r, http.StatusOK)
	return nil
}

func DeleteCity(w http.ResponseWriter, r *http.Request) error { //удаление информации о городе по указанному id;
	var p delRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	req := &p
	id, _ := strconv.Atoi(req.Source)
	if _, ok := data.Storage[uint(id)]; !ok {
		return errors.New("no such city")
	}
	w.Write([]byte("There is no city in map as: " + data.Storage[uint(id)].Name + "\n"))
	delete(data.Storage, uint(id))
	render.Status(r, http.StatusOK)
	return nil
}
func GetCityby(w http.ResponseWriter, r *http.Request) error {
	params := r.URL.Query()
	if len(params) != 1 { // Фильтр если послали запрос с несколькими параметрами
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неправильно сформирован запрос"))
		return nil
	}
	for k, v := range params {

		switch k {
		case "reg": // получение списка городов по указанному региону;
			output := make([]byte, 0)
			for _, va := range data.Storage {
				if va.Region == v[0] {
					Fr, _ := json.Marshal(va)
					output = append(output, Fr...)
				}
			}
			if len(output) > 0 { // если ответ пустой то отдаем 404
				w.Write([]byte(output))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Не найдено"))
				return nil
			}
		case "dist": // получение списка городов по указанному округу;
			output := make([]byte, 0)
			for _, va := range data.Storage {
				if va.District == v[0] {
					Fr, _ := json.Marshal(va)
					output = append(output, Fr...)

				}
			}
			if len(output) > 0 { // если ответ пустой то отдаем 404
				w.Write([]byte(output))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Не найдено"))
				return nil
			}

		case "id": // получение информации о городе по его id;
			id, _ := strconv.Atoi(v[0])
			if _, ok := data.Storage[uint(id)]; ok {
				Fr, _ := json.Marshal(data.Storage[uint(id)])
				w.Write([]byte(Fr))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("empty or invalid id\n"))
				return nil
			}
		case "pop": // получения списка городов по указанному диапазону численности населения;
			for _, va := range data.Storage {
				if len(v) != 2 {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Необходимо два параметра pop"))
					return nil
				}
				pop1, _ := strconv.Atoi(v[0])
				pop2, _ := strconv.Atoi(v[1])
				if pop1 > pop2 {
					pop1, pop2 = pop2, pop1
				}
				if va.Population >= uint(pop1) && va.Population <= uint(pop2) {
					Fr, _ := json.Marshal(va)
					w.Write([]byte(Fr))
				}
			}
		case "year": // получения списка городов по указанному диапазону года основания.
			for _, va := range data.Storage {
				if len(v) != 2 {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Необходимо два параметра year"))
					return nil
				}
				year1, _ := strconv.Atoi(v[0])
				year2, _ := strconv.Atoi(v[1])
				if year1 > year2 {
					year1, year2 = year2, year1
				}
				if va.Foundation >= uint16(year1) && va.Foundation <= uint16(year2) {
					Fr, _ := json.Marshal(va)
					w.Write([]byte(Fr))
				}
			}
		default:
			{
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Неправильно сформирован запрос"))
				return nil
			}
		}

	}

	render.Status(r, http.StatusOK)
	return nil
}
